package dto

import (
	"time"
)

// MarkAttendanceRequest represents the request to mark attendance
type MarkAttendanceRequest struct {
	EmployeeID     uint64    `json:"employee_id" validate:"required"`
	AttendanceDate time.Time `json:"attendance_date" validate:"required"`
	CheckInTime    time.Time `json:"check_in_time,omitempty"`
	CheckOutTime   time.Time `json:"check_out_time,omitempty"`
	Status         string    `json:"status" validate:"required,oneof=present absent late half_day"`
	Notes          string    `json:"notes,omitempty" validate:"omitempty,max=500"`
	MarkedByAdmin  bool      `json:"marked_by_admin,omitempty"`
}

// UpdateAttendanceRequest represents the request to update attendance
type UpdateAttendanceRequest struct {
	CheckInTime   time.Time `json:"check_in_time,omitempty"`
	CheckOutTime  time.Time `json:"check_out_time,omitempty"`
	Status        string    `json:"status,omitempty" validate:"omitempty,oneof=present absent late half_day"`
	Notes         string    `json:"notes,omitempty" validate:"omitempty,max=500"`
	MarkedByAdmin *bool     `json:"marked_by_admin,omitempty"`
}

// AttendanceResponse represents the attendance response structure
type AttendanceResponse struct {
	ID             uint64    `json:"id"`
	EmployeeID     uint64    `json:"employee_id"`
	EmployeeName   string    `json:"employee_name"`
	EmployeeCode   string    `json:"employee_code"`
	AttendanceDate time.Time `json:"attendance_date"`
	CheckInTime    time.Time `json:"check_in_time,omitempty"`
	CheckOutTime   time.Time `json:"check_out_time,omitempty"`
	Status         string    `json:"status"`
	IsWeekend      bool      `json:"is_weekend"`
	Notes          string    `json:"notes,omitempty"`
	MarkedByAdmin  bool      `json:"marked_by_admin"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedAt     time.Time `json:"modified_at"`
}

// AttendanceListResponse represents the response for attendance list
type AttendanceListResponse struct {
	Attendances []AttendanceResponse `json:"attendances"`
	Total       int                  `json:"total"`
	Page        int                  `json:"page"`
	Limit       int                  `json:"limit"`
}

// AttendanceQueryParams represents query parameters for attendance list
type AttendanceQueryParams struct {
	Page           int       `query:"page" validate:"omitempty,min=1"`
	Limit          int       `query:"limit" validate:"omitempty,min=1,max=100"`
	EmployeeID     uint64    `query:"employee_id" validate:"omitempty"`
	StartDate      time.Time `query:"start_date" validate:"omitempty"`
	EndDate        time.Time `query:"end_date" validate:"omitempty"`
	Status         string    `query:"status" validate:"omitempty,oneof=present absent late half_day"`
	IncludeWeekend *bool     `query:"include_weekend"`
}

// DailyAttendanceSummary represents daily attendance summary
type DailyAttendanceSummary struct {
	Date           time.Time `json:"date"`
	TotalEmployees int       `json:"total_employees"`
	PresentCount   int       `json:"present_count"`
	AbsentCount    int       `json:"absent_count"`
	LateCount      int       `json:"late_count"`
	HalfDayCount   int       `json:"half_day_count"`
	PresentPercent float64   `json:"present_percent"`
	AbsentPercent  float64   `json:"absent_percent"`
}

// BulkMarkAttendanceRequest represents bulk attendance marking
type BulkMarkAttendanceRequest struct {
	AttendanceDate time.Time            `json:"attendance_date" validate:"required"`
	Attendances    []BulkAttendanceItem `json:"attendances" validate:"required,dive"`
}

// BulkAttendanceItem represents individual attendance in bulk request
type BulkAttendanceItem struct {
	EmployeeID   uint64    `json:"employee_id" validate:"required"`
	CheckInTime  time.Time `json:"check_in_time,omitempty"`
	CheckOutTime time.Time `json:"check_out_time,omitempty"`
	Status       string    `json:"status" validate:"required,oneof=present absent late half_day"`
	Notes        string    `json:"notes,omitempty" validate:"omitempty,max=500"`
}
