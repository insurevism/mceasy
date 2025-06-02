package dto

import (
	"time"
)

// CalculateSalaryRequest represents the request to calculate monthly salary
type CalculateSalaryRequest struct {
	EmployeeID         uint64    `json:"employee_id" validate:"required"`
	CalculationMonth   time.Time `json:"calculation_month" validate:"required"`
	OverrideBaseSalary *float64  `json:"override_base_salary,omitempty" validate:"omitempty,min=0"`
}

// BulkCalculateSalaryRequest represents bulk salary calculation request
type BulkCalculateSalaryRequest struct {
	CalculationMonth time.Time `json:"calculation_month" validate:"required"`
	EmployeeIDs      []uint64  `json:"employee_ids,omitempty"`
}

// UpdateSalaryCalculationRequest represents the request to update salary calculation
type UpdateSalaryCalculationRequest struct {
	BaseSalary         *float64 `json:"base_salary,omitempty" validate:"omitempty,min=0"`
	TotalWorkingDays   *int     `json:"total_working_days,omitempty" validate:"omitempty,min=0"`
	AbsentDays         *int     `json:"absent_days,omitempty" validate:"omitempty,min=0"`
	PresentDays        *int     `json:"present_days,omitempty" validate:"omitempty,min=0"`
	FinalSalary        *float64 `json:"final_salary,omitempty" validate:"omitempty,min=0"`
	DeductionAmount    *float64 `json:"deduction_amount,omitempty" validate:"omitempty,min=0"`
	CalculationFormula *string  `json:"calculation_formula,omitempty" validate:"omitempty,max=1000"`
}

// SalaryCalculationResponse represents the salary calculation response structure
type SalaryCalculationResponse struct {
	ID                 uint64    `json:"id"`
	EmployeeID         string    `json:"employee_id"`
	EmployeeName       string    `json:"employee_name"`
	CalculationMonth   time.Time `json:"calculation_month"`
	BaseSalary         float64   `json:"base_salary"`
	TotalWorkingDays   int       `json:"total_working_days"`
	AbsentDays         int       `json:"absent_days"`
	PresentDays        int       `json:"present_days"`
	FinalSalary        float64   `json:"final_salary"`
	DeductionAmount    float64   `json:"deduction_amount"`
	CalculationFormula string    `json:"calculation_formula"`
	CreatedAt          time.Time `json:"created_at"`
	ModifiedAt         time.Time `json:"modified_at"`
}

// SalaryCalculationListResponse represents the response for salary calculation list
type SalaryCalculationListResponse struct {
	SalaryCalculations []SalaryCalculationResponse `json:"salary_calculations"`
	Total              int                         `json:"total"`
	Page               int                         `json:"page"`
	Limit              int                         `json:"limit"`
}

// SalaryCalculationQueryParams represents query parameters for salary calculation list
type SalaryCalculationQueryParams struct {
	Page             int       `query:"page" validate:"omitempty,min=1"`
	Limit            int       `query:"limit" validate:"omitempty,min=1,max=100"`
	EmployeeID       uint64    `query:"employee_id" validate:"omitempty"`
	CalculationMonth time.Time `query:"calculation_month" validate:"omitempty"`
	StartMonth       time.Time `query:"start_month" validate:"omitempty"`
	EndMonth         time.Time `query:"end_month" validate:"omitempty"`
}

// MonthlySalarySummary represents monthly salary summary
type MonthlySalarySummary struct {
	CalculationMonth  time.Time `json:"calculation_month"`
	TotalEmployees    int       `json:"total_employees"`
	CalculatedCount   int       `json:"calculated_count"`
	PendingCount      int       `json:"pending_count"`
	TotalBaseSalary   float64   `json:"total_base_salary"`
	TotalFinalSalary  float64   `json:"total_final_salary"`
	TotalDeductions   float64   `json:"total_deductions"`
	AverageAttendance float64   `json:"average_attendance_percentage"`
}

// EmployeeSalarySummary represents salary summary for an employee across months
type EmployeeSalarySummary struct {
	EmployeeID        uint64                      `json:"employee_id"`
	EmployeeName      string                      `json:"employee_name"`
	EmployeeCode      string                      `json:"employee_code"`
	TotalMonths       int                         `json:"total_months"`
	TotalBaseSalary   float64                     `json:"total_base_salary"`
	TotalFinalSalary  float64                     `json:"total_final_salary"`
	TotalDeductions   float64                     `json:"total_deductions"`
	AverageAttendance float64                     `json:"average_attendance_percentage"`
	MonthlySalaries   []SalaryCalculationResponse `json:"monthly_salaries"`
}
