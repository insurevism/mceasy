package dto

import (
	"fmt"
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

// MarkAttendanceRequestFlexible represents a flexible input for attendance marking
type MarkAttendanceRequestFlexible struct {
	EmployeeID     uint64 `json:"employee_id" validate:"required"`
	AttendanceDate string `json:"attendance_date" validate:"required"`
	CheckInTime    string `json:"check_in_time,omitempty"`
	CheckOutTime   string `json:"check_out_time,omitempty"`
	Status         string `json:"status" validate:"required,oneof=present absent late half_day"`
	Notes          string `json:"notes,omitempty" validate:"omitempty,max=500"`
	MarkedByAdmin  bool   `json:"marked_by_admin,omitempty"`
}

// ToMarkAttendanceRequest converts flexible input to standard DTO
func (f *MarkAttendanceRequestFlexible) ToMarkAttendanceRequest() (*MarkAttendanceRequest, error) {
	// Parse attendance date
	attendanceDate, err := parseFlexibleDate(f.AttendanceDate)
	if err != nil {
		return nil, fmt.Errorf("invalid attendance date: %w", err)
	}

	// Parse check-in time if provided
	var checkInTime time.Time
	if f.CheckInTime != "" {
		checkInTime, err = parseFlexibleTime(f.CheckInTime, attendanceDate)
		if err != nil {
			return nil, fmt.Errorf("invalid check-in time: %w", err)
		}
	}

	// Parse check-out time if provided
	var checkOutTime time.Time
	if f.CheckOutTime != "" {
		checkOutTime, err = parseFlexibleTime(f.CheckOutTime, attendanceDate)
		if err != nil {
			return nil, fmt.Errorf("invalid check-out time: %w", err)
		}
	}

	return &MarkAttendanceRequest{
		EmployeeID:     f.EmployeeID,
		AttendanceDate: attendanceDate,
		CheckInTime:    checkInTime,
		CheckOutTime:   checkOutTime,
		Status:         f.Status,
		Notes:          f.Notes,
		MarkedByAdmin:  f.MarkedByAdmin,
	}, nil
}

// parseFlexibleDate parses date string in various formats
func parseFlexibleDate(dateStr string) (time.Time, error) {
	// Try different date formats
	formats := []string{
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			// Return just the date part (start of day)
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s, expected YYYY-MM-DD", dateStr)
}

// parseFlexibleTime parses time string in various formats
func parseFlexibleTime(timeStr string, baseDate time.Time) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	// Try different time formats
	formats := []string{
		"15:04:05",
		"15:04",
		"03:04:05 PM",
		"03:04 PM",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			// Combine with the base date
			return time.Date(
				baseDate.Year(), baseDate.Month(), baseDate.Day(),
				t.Hour(), t.Minute(), t.Second(), 0, time.UTC,
			), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format: %s, expected HH:MM:SS or HH:MM", timeStr)
}

// UpdateAttendanceRequestFlexible represents a flexible input for updating attendance
type UpdateAttendanceRequestFlexible struct {
	CheckInTime   string `json:"check_in_time,omitempty"`
	CheckOutTime  string `json:"check_out_time,omitempty"`
	Status        string `json:"status,omitempty" validate:"omitempty,oneof=present absent late half_day"`
	Notes         string `json:"notes,omitempty" validate:"omitempty,max=500"`
	MarkedByAdmin *bool  `json:"marked_by_admin,omitempty"`
}

// ToUpdateAttendanceRequest converts flexible input to standard DTO
func (f *UpdateAttendanceRequestFlexible) ToUpdateAttendanceRequest() (*UpdateAttendanceRequest, error) {
	// We need a base date for time parsing - use today as default
	baseDate := time.Now()

	// Parse check-in time if provided
	var checkInTime time.Time
	if f.CheckInTime != "" {
		var err error
		checkInTime, err = parseFlexibleTime(f.CheckInTime, baseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid check-in time: %w", err)
		}
	}

	// Parse check-out time if provided
	var checkOutTime time.Time
	if f.CheckOutTime != "" {
		var err error
		checkOutTime, err = parseFlexibleTime(f.CheckOutTime, baseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid check-out time: %w", err)
		}
	}

	return &UpdateAttendanceRequest{
		CheckInTime:   checkInTime,
		CheckOutTime:  checkOutTime,
		Status:        f.Status,
		Notes:         f.Notes,
		MarkedByAdmin: f.MarkedByAdmin,
	}, nil
}
