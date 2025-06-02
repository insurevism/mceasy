package repository

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/ent/employee"
	"mceasy/internal/applications/employee/dto"
)

// EmployeeRepository defines the interface for employee data operations
type EmployeeRepository interface {
	Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*ent.Employee, error)
	GetByID(ctx context.Context, id uint64) (*ent.Employee, error)
	GetByEmployeeID(ctx context.Context, employeeID string) (*ent.Employee, error)
	Update(ctx context.Context, id uint64, req *dto.UpdateEmployeeRequest) (*ent.Employee, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, params *dto.EmployeeQueryParams) ([]*ent.Employee, int, error)
	GetActiveEmployees(ctx context.Context) ([]*ent.Employee, error)
}

// EmployeeRepositoryImpl implements the EmployeeRepository interface
type EmployeeRepositoryImpl struct {
	client *ent.Client
}

// NewEmployeeRepository creates a new employee repository instance
func NewEmployeeRepository(client *ent.Client) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{
		client: client,
	}
}

// Create creates a new employee
func (r *EmployeeRepositoryImpl) Create(ctx context.Context, req *dto.CreateEmployeeRequest) (*ent.Employee, error) {
	query := r.client.Employee.Create().
		SetFullName(req.FullName).
		SetEmail(req.Email).
		SetHireDate(req.HireDate)

	if req.Phone != "" {
		query = query.SetPhone(req.Phone)
	}
	if req.Position != "" {
		query = query.SetPosition(req.Position)
	}
	if req.Department != "" {
		query = query.SetDepartment(req.Department)
	}
	if req.BaseSalary > 0 {
		query = query.SetBaseSalary(req.BaseSalary)
	}

	return query.Save(ctx)
}

// GetByID retrieves an employee by ID
func (r *EmployeeRepositoryImpl) GetByID(ctx context.Context, id uint64) (*ent.Employee, error) {
	return r.client.Employee.
		Query().
		Where(employee.ID(id)).
		Where(employee.DeletedAtIsNil()).
		First(ctx)
}

// GetByEmployeeID retrieves an employee by employee ID
func (r *EmployeeRepositoryImpl) GetByEmployeeID(ctx context.Context, employeeID string) (*ent.Employee, error) {
	return r.client.Employee.
		Query().
		Where(employee.EmployeeID(employeeID)).
		Where(employee.DeletedAtIsNil()).
		First(ctx)
}

// Update updates an employee
func (r *EmployeeRepositoryImpl) Update(ctx context.Context, id uint64, req *dto.UpdateEmployeeRequest) (*ent.Employee, error) {
	query := r.client.Employee.UpdateOneID(id)

	if req.FullName != "" {
		query = query.SetFullName(req.FullName)
	}
	if req.Email != "" {
		query = query.SetEmail(req.Email)
	}
	if req.Phone != "" {
		query = query.SetPhone(req.Phone)
	}
	if req.Position != "" {
		query = query.SetPosition(req.Position)
	}
	if req.Department != "" {
		query = query.SetDepartment(req.Department)
	}
	if !req.HireDate.IsZero() {
		query = query.SetHireDate(req.HireDate)
	}
	if req.BaseSalary > 0 {
		query = query.SetBaseSalary(req.BaseSalary)
	}
	if req.IsActive != nil {
		query = query.SetIsActive(*req.IsActive)
	}

	return query.Save(ctx)
}

// Delete soft deletes an employee
func (r *EmployeeRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.client.Employee.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)
}

// List retrieves employees with pagination and filtering
func (r *EmployeeRepositoryImpl) List(ctx context.Context, params *dto.EmployeeQueryParams) ([]*ent.Employee, int, error) {
	query := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil())

	// Apply filters
	if params.Search != "" {
		query = query.Where(
			employee.Or(
				employee.FullNameContainsFold(params.Search),
				employee.EmailContainsFold(params.Search),
				employee.EmployeeIDContainsFold(params.Search),
			),
		)
	}

	if params.Department != "" {
		query = query.Where(employee.DepartmentEQ(params.Department))
	}

	if params.IsActive != nil {
		query = query.Where(employee.IsActiveEQ(*params.IsActive))
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count employees: %w", err)
	}

	// Apply pagination
	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	// Order by created_at desc
	query = query.Order(ent.Desc(employee.FieldCreatedAt))

	employees, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch employees: %w", err)
	}

	return employees, total, nil
}

// GetActiveEmployees retrieves all active employees
func (r *EmployeeRepositoryImpl) GetActiveEmployees(ctx context.Context) ([]*ent.Employee, error) {
	return r.client.Employee.
		Query().
		Where(employee.IsActiveEQ(true)).
		Where(employee.DeletedAtIsNil()).
		Order(ent.Asc(employee.FieldEmployeeID)).
		All(ctx)
}
