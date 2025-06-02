package service

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/internal/applications/attendance/dto"
	"mceasy/internal/applications/attendance/repository"
	"mceasy/internal/component/cache"
	"mceasy/internal/component/transaction"
)

// AttendanceService defines the interface for attendance business logic
type AttendanceService interface {
	MarkAttendance(ctx context.Context, req *dto.MarkAttendanceRequest) (*dto.AttendanceResponse, error)
	GetAttendanceByID(ctx context.Context, id uint64) (*dto.AttendanceResponse, error)
	GetAttendanceByEmployeeAndDate(ctx context.Context, employeeID uint64, date time.Time) (*dto.AttendanceResponse, error)
	UpdateAttendance(ctx context.Context, id uint64, req *dto.UpdateAttendanceRequest) (*dto.AttendanceResponse, error)
	DeleteAttendance(ctx context.Context, id uint64) error
	ListAttendance(ctx context.Context, params *dto.AttendanceQueryParams) (*dto.AttendanceListResponse, error)
	GetTodayAttendance(ctx context.Context) ([]dto.AttendanceResponse, error)
	GetDailyAttendanceSummary(ctx context.Context, date time.Time) (*dto.DailyAttendanceSummary, error)
	CheckInEmployee(ctx context.Context, employeeID uint64, checkInTime time.Time) (*dto.AttendanceResponse, error)
	CheckOutEmployee(ctx context.Context, employeeID uint64, checkOutTime time.Time) (*dto.AttendanceResponse, error)
	AutoMarkAbsentEmployees(ctx context.Context) error
}

// AttendanceServiceImpl implements the AttendanceService interface
type AttendanceServiceImpl struct {
	attendanceRepo repository.AttendanceRepository
	cache          cache.Cache
	trx            transaction.Trx
}

// NewAttendanceService creates a new attendance service instance
func NewAttendanceService(
	attendanceRepo repository.AttendanceRepository,
	cache cache.Cache,
	trx transaction.Trx,
) *AttendanceServiceImpl {
	return &AttendanceServiceImpl{
		attendanceRepo: attendanceRepo,
		cache:          cache,
		trx:            trx,
	}
}

// MarkAttendance marks attendance for an employee
func (s *AttendanceServiceImpl) MarkAttendance(ctx context.Context, req *dto.MarkAttendanceRequest) (*dto.AttendanceResponse, error) {
	// Validate attendance date (not future date)
	if req.AttendanceDate.After(time.Now()) {
		return nil, fmt.Errorf("cannot mark attendance for future dates")
	}

	// Check if it's a weekend and adjust status if needed
	if isWeekend(req.AttendanceDate) && req.Status == "present" {
		return nil, fmt.Errorf("cannot mark present on weekends")
	}

	// Validate check-in and check-out times
	if !req.CheckInTime.IsZero() && !req.CheckOutTime.IsZero() {
		if req.CheckOutTime.Before(req.CheckInTime) {
			return nil, fmt.Errorf("check-out time cannot be before check-in time")
		}
	}

	// Determine if employee is late (after 9:00 AM)
	if req.Status == "present" && !req.CheckInTime.IsZero() {
		nineAM := time.Date(req.AttendanceDate.Year(), req.AttendanceDate.Month(), req.AttendanceDate.Day(), 9, 0, 0, 0, req.AttendanceDate.Location())
		if req.CheckInTime.After(nineAM) {
			req.Status = "late"
		}
	}

	attendance, err := s.attendanceRepo.MarkAttendance(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to mark attendance: %w", err)
	}

	return s.mapToAttendanceResponse(attendance), nil
}

// GetAttendanceByID retrieves an attendance record by ID
func (s *AttendanceServiceImpl) GetAttendanceByID(ctx context.Context, id uint64) (*dto.AttendanceResponse, error) {
	attendance, err := s.attendanceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("attendance record not found: %w", err)
	}

	return s.mapToAttendanceResponse(attendance), nil
}

// GetAttendanceByEmployeeAndDate retrieves attendance by employee and date
func (s *AttendanceServiceImpl) GetAttendanceByEmployeeAndDate(ctx context.Context, employeeID uint64, date time.Time) (*dto.AttendanceResponse, error) {
	attendance, err := s.attendanceRepo.GetByEmployeeAndDate(ctx, employeeID, date)
	if err != nil {
		return nil, fmt.Errorf("attendance record not found: %w", err)
	}

	return s.mapToAttendanceResponse(attendance), nil
}

// UpdateAttendance updates an attendance record
func (s *AttendanceServiceImpl) UpdateAttendance(ctx context.Context, id uint64, req *dto.UpdateAttendanceRequest) (*dto.AttendanceResponse, error) {
	// Check if attendance exists
	_, err := s.attendanceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("attendance record not found: %w", err)
	}

	// Validate check-in and check-out times
	if !req.CheckInTime.IsZero() && !req.CheckOutTime.IsZero() {
		if req.CheckOutTime.Before(req.CheckInTime) {
			return nil, fmt.Errorf("check-out time cannot be before check-in time")
		}
	}

	attendance, err := s.attendanceRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update attendance: %w", err)
	}

	return s.mapToAttendanceResponse(attendance), nil
}

// DeleteAttendance soft deletes an attendance record
func (s *AttendanceServiceImpl) DeleteAttendance(ctx context.Context, id uint64) error {
	// Check if attendance exists
	_, err := s.attendanceRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("attendance record not found: %w", err)
	}

	err = s.attendanceRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete attendance: %w", err)
	}

	return nil
}

// ListAttendance retrieves attendance records with pagination and filtering
func (s *AttendanceServiceImpl) ListAttendance(ctx context.Context, params *dto.AttendanceQueryParams) (*dto.AttendanceListResponse, error) {
	// Set default pagination if not provided
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	attendances, total, err := s.attendanceRepo.List(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendance records: %w", err)
	}

	attendanceResponses := make([]dto.AttendanceResponse, len(attendances))
	for i, attendance := range attendances {
		attendanceResponses[i] = *s.mapToAttendanceResponse(attendance)
	}

	return &dto.AttendanceListResponse{
		Attendances: attendanceResponses,
		Total:       total,
		Page:        params.Page,
		Limit:       params.Limit,
	}, nil
}

// GetTodayAttendance retrieves today's attendance records
func (s *AttendanceServiceImpl) GetTodayAttendance(ctx context.Context) ([]dto.AttendanceResponse, error) {
	attendances, err := s.attendanceRepo.GetTodayAttendance(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's attendance: %w", err)
	}

	attendanceResponses := make([]dto.AttendanceResponse, len(attendances))
	for i, attendance := range attendances {
		attendanceResponses[i] = *s.mapToAttendanceResponse(attendance)
	}

	return attendanceResponses, nil
}

// GetDailyAttendanceSummary retrieves daily attendance summary
func (s *AttendanceServiceImpl) GetDailyAttendanceSummary(ctx context.Context, date time.Time) (*dto.DailyAttendanceSummary, error) {
	summary, err := s.attendanceRepo.GetDailyAttendanceSummary(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily attendance summary: %w", err)
	}

	return summary, nil
}

// CheckInEmployee marks an employee as present with check-in time
func (s *AttendanceServiceImpl) CheckInEmployee(ctx context.Context, employeeID uint64, checkInTime time.Time) (*dto.AttendanceResponse, error) {
	today := time.Now()
	attendanceDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	// Determine status based on check-in time
	status := "present"
	nineAM := time.Date(today.Year(), today.Month(), today.Day(), 9, 0, 0, 0, today.Location())
	if checkInTime.After(nineAM) {
		status = "late"
	}

	req := &dto.MarkAttendanceRequest{
		EmployeeID:     employeeID,
		AttendanceDate: attendanceDate,
		CheckInTime:    checkInTime,
		Status:         status,
		MarkedByAdmin:  false,
	}

	return s.MarkAttendance(ctx, req)
}

// CheckOutEmployee updates check-out time for an employee
func (s *AttendanceServiceImpl) CheckOutEmployee(ctx context.Context, employeeID uint64, checkOutTime time.Time) (*dto.AttendanceResponse, error) {
	today := time.Now()
	attendanceDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	// Get existing attendance record
	attendance, err := s.attendanceRepo.GetByEmployeeAndDate(ctx, employeeID, attendanceDate)
	if err != nil {
		return nil, fmt.Errorf("no check-in record found for today")
	}

	req := &dto.UpdateAttendanceRequest{
		CheckOutTime: checkOutTime,
	}

	return s.UpdateAttendance(ctx, attendance.ID, req)
}

// AutoMarkAbsentEmployees automatically marks employees as absent if not present by 9:00 AM
func (s *AttendanceServiceImpl) AutoMarkAbsentEmployees(ctx context.Context) error {
	today := time.Now()

	// Only run on weekdays
	if isWeekend(today) {
		return nil
	}

	// Only run after 9:00 AM
	nineAM := time.Date(today.Year(), today.Month(), today.Day(), 9, 0, 0, 0, today.Location())
	if today.Before(nineAM) {
		return nil
	}

	// This would typically be called by a scheduled job
	// Implementation would get all active employees and mark absent those without attendance records
	return nil
}

// mapToAttendanceResponse maps an ent.Attendance to dto.AttendanceResponse
func (s *AttendanceServiceImpl) mapToAttendanceResponse(attendance *ent.Attendance) *dto.AttendanceResponse {
	response := &dto.AttendanceResponse{
		ID:             attendance.ID,
		EmployeeID:     attendance.EmployeeID,
		AttendanceDate: attendance.AttendanceDate,
		CheckInTime:    attendance.CheckInTime,
		CheckOutTime:   attendance.CheckOutTime,
		Status:         string(attendance.Status),
		IsWeekend:      attendance.IsWeekend,
		Notes:          attendance.Notes,
		MarkedByAdmin:  attendance.MarkedByAdmin,
		CreatedAt:      attendance.CreatedAt,
		ModifiedAt:     attendance.ModifiedAt,
	}

	// Add employee information if available
	if attendance.Edges.Employee != nil {
		response.EmployeeName = attendance.Edges.Employee.FullName
		response.EmployeeCode = attendance.Edges.Employee.EmployeeID
	}

	return response
}

// isWeekend checks if the given date is a weekend
func isWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// parseDate parses date string in format "2006-01-02" or "2006-01-02T15:04:05Z"
func parseDate(dateStr string) (time.Time, error) {
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
			return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s, expected YYYY-MM-DD", dateStr)
}

// parseTime parses time string in format "15:04:05" or "15:04"
func parseTime(timeStr string, baseDate time.Time) (time.Time, error) {
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
				t.Hour(), t.Minute(), t.Second(), 0, baseDate.Location(),
			), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format: %s, expected HH:MM:SS or HH:MM", timeStr)
}

// formatTimeString formats time to string, returns empty string for zero time
func formatTimeString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("15:04:05")
}
