package service

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/internal/applications/salary/dto"
	"mceasy/internal/applications/salary/repository"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"
)

// SalaryService defines the interface for salary calculation business logic
type SalaryService interface {
	CalculateSalary(ctx context.Context, req *dto.CalculateSalaryRequest) (*dto.SalaryCalculationResponse, error)
	GetSalaryCalculationByID(ctx context.Context, id uint64) (*dto.SalaryCalculationResponse, error)
	GetSalaryCalculationByEmployeeAndMonth(ctx context.Context, employeeID uint64, month time.Time) (*dto.SalaryCalculationResponse, error)
	UpdateSalaryCalculation(ctx context.Context, id uint64, req *dto.UpdateSalaryCalculationRequest) (*dto.SalaryCalculationResponse, error)
	DeleteSalaryCalculation(ctx context.Context, id uint64) error
	ListSalaryCalculations(ctx context.Context, params *dto.SalaryCalculationQueryParams) (*dto.SalaryCalculationListResponse, error)
	GetMonthlySalarySummary(ctx context.Context, month time.Time) (*dto.MonthlySalarySummary, error)
	GetEmployeeSalarySummary(ctx context.Context, employeeID uint64, startMonth, endMonth time.Time) (*dto.EmployeeSalarySummary, error)
	BulkCalculateSalary(ctx context.Context, req *dto.BulkCalculateSalaryRequest) ([]dto.SalaryCalculationResponse, error)
}

// SalaryServiceImpl implements the SalaryService interface
type SalaryServiceImpl struct {
	salaryRepo repository.SalaryRepository
	cache      cache.Cache
	trx        transaction.Trx
}

// NewSalaryService creates a new salary service instance
func NewSalaryService(
	salaryRepo repository.SalaryRepository,
	cache cache.Cache,
	trx transaction.Trx,
) *SalaryServiceImpl {
	return &SalaryServiceImpl{
		salaryRepo: salaryRepo,
		cache:      cache,
		trx:        trx,
	}
}

// CalculateSalary calculates monthly salary for an employee
func (s *SalaryServiceImpl) CalculateSalary(ctx context.Context, req *dto.CalculateSalaryRequest) (*dto.SalaryCalculationResponse, error) {
	// Validate calculation month (not future month)
	currentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	reqMonth := time.Date(req.CalculationMonth.Year(), req.CalculationMonth.Month(), 1, 0, 0, 0, 0, req.CalculationMonth.Location())

	if reqMonth.After(currentMonth) {
		return nil, fmt.Errorf("cannot calculate salary for future months")
	}

	// Validate override base salary if provided
	if req.OverrideBaseSalary != nil && *req.OverrideBaseSalary < 0 {
		return nil, fmt.Errorf("base salary cannot be negative")
	}

	calculation, err := s.salaryRepo.CalculateSalary(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate salary: %w", err)
	}

	return s.mapToSalaryCalculationResponse(calculation), nil
}

// GetSalaryCalculationByID retrieves a salary calculation by ID
func (s *SalaryServiceImpl) GetSalaryCalculationByID(ctx context.Context, id uint64) (*dto.SalaryCalculationResponse, error) {
	calculation, err := s.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("salary calculation not found: %w", err)
	}

	return s.mapToSalaryCalculationResponse(calculation), nil
}

// GetSalaryCalculationByEmployeeAndMonth retrieves salary calculation by employee and month
func (s *SalaryServiceImpl) GetSalaryCalculationByEmployeeAndMonth(ctx context.Context, employeeID uint64, month time.Time) (*dto.SalaryCalculationResponse, error) {
	calculation, err := s.salaryRepo.GetByEmployeeAndMonth(ctx, employeeID, month)
	if err != nil {
		return nil, fmt.Errorf("salary calculation not found: %w", err)
	}

	return s.mapToSalaryCalculationResponse(calculation), nil
}

// UpdateSalaryCalculation updates a salary calculation
func (s *SalaryServiceImpl) UpdateSalaryCalculation(ctx context.Context, id uint64, req *dto.UpdateSalaryCalculationRequest) (*dto.SalaryCalculationResponse, error) {
	// Validate input values
	if req.BaseSalary != nil && *req.BaseSalary < 0 {
		return nil, fmt.Errorf("base salary cannot be negative")
	}
	if req.TotalWorkingDays != nil && *req.TotalWorkingDays < 0 {
		return nil, fmt.Errorf("total working days cannot be negative")
	}
	if req.AbsentDays != nil && *req.AbsentDays < 0 {
		return nil, fmt.Errorf("absent days cannot be negative")
	}
	if req.PresentDays != nil && *req.PresentDays < 0 {
		return nil, fmt.Errorf("present days cannot be negative")
	}
	if req.FinalSalary != nil && *req.FinalSalary < 0 {
		return nil, fmt.Errorf("final salary cannot be negative")
	}
	if req.DeductionAmount != nil && *req.DeductionAmount < 0 {
		return nil, fmt.Errorf("deduction amount cannot be negative")
	}

	// Check if salary calculation exists
	_, err := s.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("salary calculation not found: %w", err)
	}

	calculation, err := s.salaryRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update salary calculation: %w", err)
	}

	return s.mapToSalaryCalculationResponse(calculation), nil
}

// DeleteSalaryCalculation soft deletes a salary calculation
func (s *SalaryServiceImpl) DeleteSalaryCalculation(ctx context.Context, id uint64) error {
	// Check if salary calculation exists
	_, err := s.salaryRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("salary calculation not found: %w", err)
	}

	err = s.salaryRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete salary calculation: %w", err)
	}

	return nil
}

// ListSalaryCalculations retrieves salary calculations with pagination and filtering
func (s *SalaryServiceImpl) ListSalaryCalculations(ctx context.Context, params *dto.SalaryCalculationQueryParams) (*dto.SalaryCalculationListResponse, error) {
	// Set default pagination if not provided
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	calculations, total, err := s.salaryRepo.List(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list salary calculations: %w", err)
	}

	calculationResponses := make([]dto.SalaryCalculationResponse, len(calculations))
	for i, calculation := range calculations {
		calculationResponses[i] = *s.mapToSalaryCalculationResponse(calculation)
	}

	return &dto.SalaryCalculationListResponse{
		SalaryCalculations: calculationResponses,
		Total:              total,
		Page:               params.Page,
		Limit:              params.Limit,
	}, nil
}

// GetMonthlySalarySummary retrieves monthly salary summary
func (s *SalaryServiceImpl) GetMonthlySalarySummary(ctx context.Context, month time.Time) (*dto.MonthlySalarySummary, error) {
	summary, err := s.salaryRepo.GetMonthlySalarySummary(ctx, month)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly salary summary: %w", err)
	}

	return summary, nil
}

// GetEmployeeSalarySummary retrieves salary summary for an employee across months
func (s *SalaryServiceImpl) GetEmployeeSalarySummary(ctx context.Context, employeeID uint64, startMonth, endMonth time.Time) (*dto.EmployeeSalarySummary, error) {
	// Validate date range
	if endMonth.Before(startMonth) {
		return nil, fmt.Errorf("end month cannot be before start month")
	}

	summary, err := s.salaryRepo.GetEmployeeSalarySummary(ctx, employeeID, startMonth, endMonth)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee salary summary: %w", err)
	}

	return summary, nil
}

// BulkCalculateSalary calculates salary for multiple employees
func (s *SalaryServiceImpl) BulkCalculateSalary(ctx context.Context, req *dto.BulkCalculateSalaryRequest) ([]dto.SalaryCalculationResponse, error) {
	// Validate calculation month (not future month)
	currentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	reqMonth := time.Date(req.CalculationMonth.Year(), req.CalculationMonth.Month(), 1, 0, 0, 0, 0, req.CalculationMonth.Location())

	if reqMonth.After(currentMonth) {
		return nil, fmt.Errorf("cannot calculate salary for future months")
	}

	calculations, err := s.salaryRepo.BulkCalculateSalary(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk calculate salary: %w", err)
	}

	calculationResponses := make([]dto.SalaryCalculationResponse, len(calculations))
	for i, calculation := range calculations {
		calculationResponses[i] = *s.mapToSalaryCalculationResponse(calculation)
	}

	return calculationResponses, nil
}

// mapToSalaryCalculationResponse converts ENT entity to response DTO
func (s *SalaryServiceImpl) mapToSalaryCalculationResponse(calculation *ent.SalaryCalculation) *dto.SalaryCalculationResponse {
	response := &dto.SalaryCalculationResponse{
		ID:                 calculation.ID,
		CalculationMonth:   calculation.CalculationMonth,
		BaseSalary:         calculation.BaseSalary,
		TotalWorkingDays:   calculation.TotalWorkingDays,
		AbsentDays:         calculation.AbsentDays,
		PresentDays:        calculation.PresentDays,
		FinalSalary:        calculation.FinalSalary,
		DeductionAmount:    calculation.DeductionAmount,
		CalculationFormula: calculation.CalculationFormula,
		CreatedAt:          calculation.CreatedAt,
		ModifiedAt:         calculation.ModifiedAt,
	}

	// Add employee information if available
	if calculation.Edges.Employee != nil {
		response.EmployeeName = calculation.Edges.Employee.FullName
		response.EmployeeID = calculation.Edges.Employee.EmployeeID
	}

	return response
}
