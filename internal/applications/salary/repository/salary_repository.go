package repository

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/ent/attendance"
	"mceasy/ent/employee"
	"mceasy/ent/salarycalculation"
	"mceasy/internal/applications/salary/dto"
)

// SalaryRepository defines the interface for salary calculation data operations
type SalaryRepository interface {
	CalculateSalary(ctx context.Context, req *dto.CalculateSalaryRequest) (*ent.SalaryCalculation, error)
	GetByID(ctx context.Context, id uint64) (*ent.SalaryCalculation, error)
	GetByEmployeeAndMonth(ctx context.Context, employeeID uint64, month time.Time) (*ent.SalaryCalculation, error)
	Update(ctx context.Context, id uint64, req *dto.UpdateSalaryCalculationRequest) (*ent.SalaryCalculation, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, params *dto.SalaryCalculationQueryParams) ([]*ent.SalaryCalculation, int, error)
	GetMonthlySalarySummary(ctx context.Context, month time.Time) (*dto.MonthlySalarySummary, error)
	GetEmployeeSalarySummary(ctx context.Context, employeeID uint64, startMonth, endMonth time.Time) (*dto.EmployeeSalarySummary, error)
	BulkCalculateSalary(ctx context.Context, req *dto.BulkCalculateSalaryRequest) ([]*ent.SalaryCalculation, error)
	GetWorkingDaysInMonth(ctx context.Context, month time.Time) (int, error)
	GetAttendanceDataForMonth(ctx context.Context, employeeID uint64, month time.Time) (presentDays, absentDays int, err error)
}

// SalaryRepositoryImpl implements the SalaryRepository interface
type SalaryRepositoryImpl struct {
	client *ent.Client
}

// NewSalaryRepository creates a new salary repository instance
func NewSalaryRepository(client *ent.Client) *SalaryRepositoryImpl {
	return &SalaryRepositoryImpl{
		client: client,
	}
}

// CalculateSalary creates or updates a salary calculation
func (r *SalaryRepositoryImpl) CalculateSalary(ctx context.Context, req *dto.CalculateSalaryRequest) (*ent.SalaryCalculation, error) {
	// Get employee data
	emp, err := r.client.Employee.
		Query().
		Where(employee.ID(req.EmployeeID)).
		Where(employee.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	// Normalize month to first day of month
	normalizedMonth := time.Date(req.CalculationMonth.Year(), req.CalculationMonth.Month(), 1, 0, 0, 0, 0, req.CalculationMonth.Location())

	// Check if calculation already exists
	existing, _ := r.GetByEmployeeAndMonth(ctx, req.EmployeeID, normalizedMonth)

	// Get working days in month
	totalWorkingDays, err := r.GetWorkingDaysInMonth(ctx, normalizedMonth)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate working days: %w", err)
	}

	// Get attendance data for the month
	presentDays, absentDays, err := r.GetAttendanceDataForMonth(ctx, req.EmployeeID, normalizedMonth)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance data: %w", err)
	}

	// Determine base salary
	baseSalary := emp.BaseSalary
	if req.OverrideBaseSalary != nil {
		baseSalary = *req.OverrideBaseSalary
	}

	// Calculate salary: proportional deduction based on absent days
	// Formula: final_salary = base_salary * (present_days / total_working_days)
	finalSalary := baseSalary
	deductionAmount := 0.0

	if totalWorkingDays > 0 {
		finalSalary = baseSalary * (float64(presentDays) / float64(totalWorkingDays))
		deductionAmount = baseSalary - finalSalary
	}

	// Format calculation formula
	calculationFormula := fmt.Sprintf("Base: %.2f, Working Days: %d, Present: %d, Absent: %d, Final: %.2f * (%d/%d) = %.2f",
		baseSalary, totalWorkingDays, presentDays, absentDays, baseSalary, presentDays, totalWorkingDays, finalSalary)

	if existing != nil {
		// Update existing calculation
		return r.client.SalaryCalculation.
			UpdateOneID(existing.ID).
			SetBaseSalary(baseSalary).
			SetTotalWorkingDays(totalWorkingDays).
			SetAbsentDays(absentDays).
			SetPresentDays(presentDays).
			SetFinalSalary(finalSalary).
			SetDeductionAmount(deductionAmount).
			SetCalculationFormula(calculationFormula).
			Save(ctx)
	}

	// Create new calculation
	return r.client.SalaryCalculation.
		Create().
		SetEmployeeID(req.EmployeeID).
		SetCalculationMonth(normalizedMonth).
		SetBaseSalary(baseSalary).
		SetTotalWorkingDays(totalWorkingDays).
		SetAbsentDays(absentDays).
		SetPresentDays(presentDays).
		SetFinalSalary(finalSalary).
		SetDeductionAmount(deductionAmount).
		SetCalculationFormula(calculationFormula).
		Save(ctx)
}

// GetByID retrieves a salary calculation by ID
func (r *SalaryRepositoryImpl) GetByID(ctx context.Context, id uint64) (*ent.SalaryCalculation, error) {
	return r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.ID(id)).
		Where(salarycalculation.DeletedAtIsNil()).
		WithEmployee().
		First(ctx)
}

// GetByEmployeeAndMonth retrieves salary calculation by employee and month
func (r *SalaryRepositoryImpl) GetByEmployeeAndMonth(ctx context.Context, employeeID uint64, month time.Time) (*ent.SalaryCalculation, error) {
	// Normalize month to first day of month
	normalizedMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())

	return r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.EmployeeID(employeeID)).
		Where(salarycalculation.CalculationMonth(normalizedMonth)).
		Where(salarycalculation.DeletedAtIsNil()).
		WithEmployee().
		First(ctx)
}

// Update updates a salary calculation record
func (r *SalaryRepositoryImpl) Update(ctx context.Context, id uint64, req *dto.UpdateSalaryCalculationRequest) (*ent.SalaryCalculation, error) {
	query := r.client.SalaryCalculation.UpdateOneID(id)

	if req.BaseSalary != nil {
		query = query.SetBaseSalary(*req.BaseSalary)
	}
	if req.TotalWorkingDays != nil {
		query = query.SetTotalWorkingDays(*req.TotalWorkingDays)
	}
	if req.AbsentDays != nil {
		query = query.SetAbsentDays(*req.AbsentDays)
	}
	if req.PresentDays != nil {
		query = query.SetPresentDays(*req.PresentDays)
	}
	if req.FinalSalary != nil {
		query = query.SetFinalSalary(*req.FinalSalary)
	}
	if req.DeductionAmount != nil {
		query = query.SetDeductionAmount(*req.DeductionAmount)
	}
	if req.CalculationFormula != nil {
		query = query.SetCalculationFormula(*req.CalculationFormula)
	}

	return query.Save(ctx)
}

// Delete soft deletes a salary calculation record
func (r *SalaryRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.client.SalaryCalculation.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)
}

// List retrieves salary calculations with pagination and filtering
func (r *SalaryRepositoryImpl) List(ctx context.Context, params *dto.SalaryCalculationQueryParams) ([]*ent.SalaryCalculation, int, error) {
	query := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.DeletedAtIsNil()).
		WithEmployee()

	// Apply filters
	if params.EmployeeID > 0 {
		query = query.Where(salarycalculation.EmployeeID(params.EmployeeID))
	}

	if !params.CalculationMonth.IsZero() {
		normalizedMonth := time.Date(params.CalculationMonth.Year(), params.CalculationMonth.Month(), 1, 0, 0, 0, 0, params.CalculationMonth.Location())
		query = query.Where(salarycalculation.CalculationMonth(normalizedMonth))
	}

	if !params.StartMonth.IsZero() {
		normalizedStartMonth := time.Date(params.StartMonth.Year(), params.StartMonth.Month(), 1, 0, 0, 0, 0, params.StartMonth.Location())
		query = query.Where(salarycalculation.CalculationMonthGTE(normalizedStartMonth))
	}

	if !params.EndMonth.IsZero() {
		normalizedEndMonth := time.Date(params.EndMonth.Year(), params.EndMonth.Month(), 1, 0, 0, 0, 0, params.EndMonth.Location())
		query = query.Where(salarycalculation.CalculationMonthLTE(normalizedEndMonth))
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count salary calculations: %w", err)
	}

	// Apply pagination
	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	// Order by calculation_month desc, then by employee_id
	query = query.Order(ent.Desc(salarycalculation.FieldCalculationMonth), ent.Asc(salarycalculation.FieldEmployeeID))

	calculations, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch salary calculations: %w", err)
	}

	return calculations, total, nil
}

// GetMonthlySalarySummary calculates monthly salary summary
func (r *SalaryRepositoryImpl) GetMonthlySalarySummary(ctx context.Context, month time.Time) (*dto.MonthlySalarySummary, error) {
	// Normalize month to first day of month
	normalizedMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())

	// Get total employees (active)
	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count total employees: %w", err)
	}

	// Get calculated salaries for the month
	calculations, err := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.CalculationMonth(normalizedMonth)).
		Where(salarycalculation.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch salary calculations: %w", err)
	}

	calculatedCount := len(calculations)
	pendingCount := totalEmployees - calculatedCount

	// Calculate totals
	var totalBaseSalary, totalFinalSalary, totalDeductions float64
	var totalAttendancePercentage float64

	for _, calc := range calculations {
		totalBaseSalary += calc.BaseSalary
		totalFinalSalary += calc.FinalSalary
		totalDeductions += calc.DeductionAmount

		if calc.TotalWorkingDays > 0 {
			attendancePercentage := float64(calc.PresentDays) / float64(calc.TotalWorkingDays) * 100
			totalAttendancePercentage += attendancePercentage
		}
	}

	averageAttendance := 0.0
	if calculatedCount > 0 {
		averageAttendance = totalAttendancePercentage / float64(calculatedCount)
	}

	return &dto.MonthlySalarySummary{
		CalculationMonth:  normalizedMonth,
		TotalEmployees:    totalEmployees,
		CalculatedCount:   calculatedCount,
		PendingCount:      pendingCount,
		TotalBaseSalary:   totalBaseSalary,
		TotalFinalSalary:  totalFinalSalary,
		TotalDeductions:   totalDeductions,
		AverageAttendance: averageAttendance,
	}, nil
}

// GetEmployeeSalarySummary retrieves salary summary for an employee across months
func (r *SalaryRepositoryImpl) GetEmployeeSalarySummary(ctx context.Context, employeeID uint64, startMonth, endMonth time.Time) (*dto.EmployeeSalarySummary, error) {
	// Get employee info
	emp, err := r.client.Employee.
		Query().
		Where(employee.ID(employeeID)).
		Where(employee.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	// Normalize months
	normalizedStartMonth := time.Date(startMonth.Year(), startMonth.Month(), 1, 0, 0, 0, 0, startMonth.Location())
	normalizedEndMonth := time.Date(endMonth.Year(), endMonth.Month(), 1, 0, 0, 0, 0, endMonth.Location())

	// Get salary calculations for the period
	calculations, err := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.EmployeeID(employeeID)).
		Where(salarycalculation.CalculationMonthGTE(normalizedStartMonth)).
		Where(salarycalculation.CalculationMonthLTE(normalizedEndMonth)).
		Where(salarycalculation.DeletedAtIsNil()).
		Order(ent.Asc(salarycalculation.FieldCalculationMonth)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch salary calculations: %w", err)
	}

	totalMonths := len(calculations)
	var totalBaseSalary, totalFinalSalary, totalDeductions float64
	var totalAttendancePercentage float64

	// Convert to response DTOs and calculate totals
	monthlySalaries := make([]dto.SalaryCalculationResponse, len(calculations))
	for i, calc := range calculations {
		monthlySalaries[i] = dto.SalaryCalculationResponse{
			ID:                 calc.ID,
			EmployeeID:         emp.EmployeeID,
			EmployeeName:       emp.FullName,
			CalculationMonth:   calc.CalculationMonth,
			BaseSalary:         calc.BaseSalary,
			TotalWorkingDays:   calc.TotalWorkingDays,
			AbsentDays:         calc.AbsentDays,
			PresentDays:        calc.PresentDays,
			FinalSalary:        calc.FinalSalary,
			DeductionAmount:    calc.DeductionAmount,
			CalculationFormula: calc.CalculationFormula,
			CreatedAt:          calc.CreatedAt,
			ModifiedAt:         calc.ModifiedAt,
		}

		totalBaseSalary += calc.BaseSalary
		totalFinalSalary += calc.FinalSalary
		totalDeductions += calc.DeductionAmount

		if calc.TotalWorkingDays > 0 {
			attendancePercentage := float64(calc.PresentDays) / float64(calc.TotalWorkingDays) * 100
			totalAttendancePercentage += attendancePercentage
		}
	}

	averageAttendance := 0.0
	if totalMonths > 0 {
		averageAttendance = totalAttendancePercentage / float64(totalMonths)
	}

	return &dto.EmployeeSalarySummary{
		EmployeeID:        employeeID,
		EmployeeName:      emp.FullName,
		EmployeeCode:      emp.EmployeeID,
		TotalMonths:       totalMonths,
		TotalBaseSalary:   totalBaseSalary,
		TotalFinalSalary:  totalFinalSalary,
		TotalDeductions:   totalDeductions,
		AverageAttendance: averageAttendance,
		MonthlySalaries:   monthlySalaries,
	}, nil
}

// BulkCalculateSalary calculates salary for multiple employees
func (r *SalaryRepositoryImpl) BulkCalculateSalary(ctx context.Context, req *dto.BulkCalculateSalaryRequest) ([]*ent.SalaryCalculation, error) {
	var employeeIDs []uint64

	if len(req.EmployeeIDs) > 0 {
		employeeIDs = req.EmployeeIDs
	} else {
		// Get all active employees
		employees, err := r.client.Employee.
			Query().
			Where(employee.DeletedAtIsNil()).
			Select(employee.FieldID).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch employees: %w", err)
		}

		for _, emp := range employees {
			employeeIDs = append(employeeIDs, emp.ID)
		}
	}

	var results []*ent.SalaryCalculation
	for _, empID := range employeeIDs {
		calcReq := &dto.CalculateSalaryRequest{
			EmployeeID:       empID,
			CalculationMonth: req.CalculationMonth,
		}

		result, err := r.CalculateSalary(ctx, calcReq)
		if err != nil {
			// Log error but continue with other employees
			continue
		}

		results = append(results, result)
	}

	return results, nil
}

// GetWorkingDaysInMonth calculates total working days in a month (excluding weekends)
func (r *SalaryRepositoryImpl) GetWorkingDaysInMonth(ctx context.Context, month time.Time) (int, error) {
	year := month.Year()
	monthNum := month.Month()

	// Get the first day and last day of the month
	firstDay := time.Date(year, monthNum, 1, 0, 0, 0, 0, month.Location())
	lastDay := firstDay.AddDate(0, 1, -1) // Last day of the month

	workingDays := 0
	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			workingDays++
		}
	}

	return workingDays, nil
}

// GetAttendanceDataForMonth retrieves attendance data for an employee in a specific month
func (r *SalaryRepositoryImpl) GetAttendanceDataForMonth(ctx context.Context, employeeID uint64, month time.Time) (presentDays, absentDays int, err error) {
	year := month.Year()
	monthNum := month.Month()

	// Get the first day and last day of the month
	firstDay := time.Date(year, monthNum, 1, 0, 0, 0, 0, month.Location())
	lastDay := firstDay.AddDate(0, 1, -1) // Last day of the month

	// Get attendance records for the month
	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.EmployeeID(employeeID)).
		Where(attendance.AttendanceDateGTE(firstDay)).
		Where(attendance.AttendanceDateLTE(lastDay)).
		Where(attendance.DeletedAtIsNil()).
		Where(attendance.IsWeekendEQ(false)). // Only working days
		All(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch attendance records: %w", err)
	}

	// Count present and absent days
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent, attendance.StatusLate, attendance.StatusHalfDay:
			presentDays++
		case attendance.StatusAbsent:
			absentDays++
		}
	}

	// Calculate working days
	totalWorkingDays, err := r.GetWorkingDaysInMonth(ctx, month)
	if err != nil {
		return 0, 0, err
	}

	// If we have fewer attendance records than working days, assume missing days are absent
	recordedDays := presentDays + absentDays
	if recordedDays < totalWorkingDays {
		absentDays += (totalWorkingDays - recordedDays)
	}

	return presentDays, absentDays, nil
}
