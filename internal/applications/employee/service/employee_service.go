package service

import (
	"context"
	"fmt"

	"mceasy/ent"
	"mceasy/internal/applications/employee/dto"
	"mceasy/internal/applications/employee/repository"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"
)

// EmployeeService defines the interface for employee business logic
type EmployeeService interface {
	CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	GetEmployeeByID(ctx context.Context, id uint64) (*dto.EmployeeResponse, error)
	GetEmployeeByEmployeeID(ctx context.Context, employeeID string) (*dto.EmployeeResponse, error)
	UpdateEmployee(ctx context.Context, id uint64, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, id uint64) error
	ListEmployees(ctx context.Context, params *dto.EmployeeQueryParams) (*dto.EmployeeListResponse, error)
}

// EmployeeServiceImpl implements the EmployeeService interface
type EmployeeServiceImpl struct {
	employeeRepo repository.EmployeeRepository
	cache        cache.Cache
	trx          transaction.Trx
}

// NewEmployeeService creates a new employee service instance
func NewEmployeeService(
	employeeRepo repository.EmployeeRepository,
	cache cache.Cache,
	trx transaction.Trx,
) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		employeeRepo: employeeRepo,
		cache:        cache,
		trx:          trx,
	}
}

// CreateEmployee creates a new employee
func (s *EmployeeServiceImpl) CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	// Check if email already exists
	existingEmployee, _ := s.employeeRepo.GetByEmployeeID(ctx, req.Email)
	if existingEmployee != nil {
		return nil, fmt.Errorf("employee with email %s already exists", req.Email)
	}

	// Set default base salary if not provided
	if req.BaseSalary == 0 {
		req.BaseSalary = 10000000.00 // Default IDR 10,000,000
	}

	employee, err := s.employeeRepo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee: %w", err)
	}

	return s.mapToEmployeeResponse(employee), nil
}

// GetEmployeeByID retrieves an employee by ID
func (s *EmployeeServiceImpl) GetEmployeeByID(ctx context.Context, id uint64) (*dto.EmployeeResponse, error) {
	employee, err := s.employeeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	return s.mapToEmployeeResponse(employee), nil
}

// GetEmployeeByEmployeeID retrieves an employee by employee ID
func (s *EmployeeServiceImpl) GetEmployeeByEmployeeID(ctx context.Context, employeeID string) (*dto.EmployeeResponse, error) {
	employee, err := s.employeeRepo.GetByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	return s.mapToEmployeeResponse(employee), nil
}

// UpdateEmployee updates an employee
func (s *EmployeeServiceImpl) UpdateEmployee(ctx context.Context, id uint64, req *dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	// Check if employee exists
	_, err := s.employeeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	// Check if email is being updated and already exists
	if req.Email != "" {
		existingEmployee, _ := s.employeeRepo.GetByEmployeeID(ctx, req.Email)
		if existingEmployee != nil && existingEmployee.ID != id {
			return nil, fmt.Errorf("employee with email %s already exists", req.Email)
		}
	}

	employee, err := s.employeeRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update employee: %w", err)
	}

	return s.mapToEmployeeResponse(employee), nil
}

// DeleteEmployee soft deletes an employee
func (s *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, id uint64) error {
	// Check if employee exists
	_, err := s.employeeRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("employee not found: %w", err)
	}

	err = s.employeeRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	return nil
}

// ListEmployees retrieves employees with pagination and filtering
func (s *EmployeeServiceImpl) ListEmployees(ctx context.Context, params *dto.EmployeeQueryParams) (*dto.EmployeeListResponse, error) {
	// Set default pagination if not provided
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	employees, total, err := s.employeeRepo.List(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list employees: %w", err)
	}

	employeeResponses := make([]dto.EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeeResponses[i] = *s.mapToEmployeeResponse(employee)
	}

	return &dto.EmployeeListResponse{
		Employees: employeeResponses,
		Total:     total,
		Page:      params.Page,
		Limit:     params.Limit,
	}, nil
}

// mapToEmployeeResponse maps an ent.Employee to dto.EmployeeResponse
func (s *EmployeeServiceImpl) mapToEmployeeResponse(employee *ent.Employee) *dto.EmployeeResponse {
	return &dto.EmployeeResponse{
		ID:         employee.ID,
		EmployeeID: employee.EmployeeID,
		FullName:   employee.FullName,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Position:   employee.Position,
		Department: employee.Department,
		HireDate:   employee.HireDate,
		BaseSalary: employee.BaseSalary,
		IsActive:   employee.IsActive,
		CreatedAt:  employee.CreatedAt,
		ModifiedAt: employee.ModifiedAt,
	}
}
