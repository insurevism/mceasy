package dto

import (
	"time"
)

// CreateEmployeeRequest represents the request to create a new employee
type CreateEmployeeRequest struct {
	FullName   string    `json:"full_name" validate:"required,min=2,max=255"`
	Email      string    `json:"email" validate:"required,email,max=255"`
	Phone      string    `json:"phone,omitempty" validate:"omitempty,max=20"`
	Position   string    `json:"position,omitempty" validate:"omitempty,max=100"`
	Department string    `json:"department,omitempty" validate:"omitempty,max=100"`
	HireDate   time.Time `json:"hire_date" validate:"required"`
	BaseSalary float64   `json:"base_salary,omitempty" validate:"omitempty,min=0"`
}

// UpdateEmployeeRequest represents the request to update an employee
type UpdateEmployeeRequest struct {
	FullName   string    `json:"full_name,omitempty" validate:"omitempty,min=2,max=255"`
	Email      string    `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Phone      string    `json:"phone,omitempty" validate:"omitempty,max=20"`
	Position   string    `json:"position,omitempty" validate:"omitempty,max=100"`
	Department string    `json:"department,omitempty" validate:"omitempty,max=100"`
	HireDate   time.Time `json:"hire_date,omitempty"`
	BaseSalary float64   `json:"base_salary,omitempty" validate:"omitempty,min=0"`
	IsActive   *bool     `json:"is_active,omitempty"`
}

// EmployeeResponse represents the employee response structure
type EmployeeResponse struct {
	ID         uint64    `json:"id"`
	EmployeeID string    `json:"employee_id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone,omitempty"`
	Position   string    `json:"position,omitempty"`
	Department string    `json:"department,omitempty"`
	HireDate   time.Time `json:"hire_date"`
	BaseSalary float64   `json:"base_salary"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// EmployeeListResponse represents the response for employee list
type EmployeeListResponse struct {
	Employees []EmployeeResponse `json:"employees"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
}

// EmployeeQueryParams represents query parameters for employee list
type EmployeeQueryParams struct {
	Page       int    `query:"page" validate:"omitempty,min=1"`
	Limit      int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search     string `query:"search" validate:"omitempty,max=255"`
	Department string `query:"department" validate:"omitempty,max=100"`
	IsActive   *bool  `query:"is_active"`
}
