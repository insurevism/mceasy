package repository

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/ent/attendance"
	"mceasy/ent/employee"
	"mceasy/internal/applications/attendance/dto"
)

// AttendanceRepository defines the interface for attendance data operations
type AttendanceRepository interface {
	MarkAttendance(ctx context.Context, req *dto.MarkAttendanceRequest) (*ent.Attendance, error)
	GetByID(ctx context.Context, id uint64) (*ent.Attendance, error)
	GetByEmployeeAndDate(ctx context.Context, employeeID uint64, date time.Time) (*ent.Attendance, error)
	Update(ctx context.Context, id uint64, req *dto.UpdateAttendanceRequest) (*ent.Attendance, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, params *dto.AttendanceQueryParams) ([]*ent.Attendance, int, error)
	GetTodayAttendance(ctx context.Context) ([]*ent.Attendance, error)
	GetDailyAttendanceSummary(ctx context.Context, date time.Time) (*dto.DailyAttendanceSummary, error)
	BulkMarkAttendance(ctx context.Context, req *dto.BulkMarkAttendanceRequest) error
	GetAttendanceByDateRange(ctx context.Context, employeeID uint64, startDate, endDate time.Time) ([]*ent.Attendance, error)
}

// AttendanceRepositoryImpl implements the AttendanceRepository interface
type AttendanceRepositoryImpl struct {
	client *ent.Client
}

// NewAttendanceRepository creates a new attendance repository instance
func NewAttendanceRepository(client *ent.Client) *AttendanceRepositoryImpl {
	return &AttendanceRepositoryImpl{
		client: client,
	}
}

// MarkAttendance creates or updates attendance record
func (r *AttendanceRepositoryImpl) MarkAttendance(ctx context.Context, req *dto.MarkAttendanceRequest) (*ent.Attendance, error) {
	// Check if attendance already exists for this employee and date
	existing, _ := r.GetByEmployeeAndDate(ctx, req.EmployeeID, req.AttendanceDate)

	if existing != nil {
		// Update existing attendance
		updateReq := &dto.UpdateAttendanceRequest{
			CheckInTime:   req.CheckInTime,
			CheckOutTime:  req.CheckOutTime,
			Status:        req.Status,
			Notes:         req.Notes,
			MarkedByAdmin: &req.MarkedByAdmin,
		}
		return r.Update(ctx, existing.ID, updateReq)
	}

	// Create new attendance record
	isWeekend := isWeekendDay(req.AttendanceDate)

	query := r.client.Attendance.Create().
		SetEmployeeID(req.EmployeeID).
		SetAttendanceDate(req.AttendanceDate).
		SetStatus(attendance.Status(req.Status)).
		SetIsWeekend(isWeekend).
		SetMarkedByAdmin(req.MarkedByAdmin)

	if !req.CheckInTime.IsZero() {
		query = query.SetCheckInTime(req.CheckInTime)
	}
	if !req.CheckOutTime.IsZero() {
		query = query.SetCheckOutTime(req.CheckOutTime)
	}
	if req.Notes != "" {
		query = query.SetNotes(req.Notes)
	}

	return query.Save(ctx)
}

// GetByID retrieves an attendance record by ID
func (r *AttendanceRepositoryImpl) GetByID(ctx context.Context, id uint64) (*ent.Attendance, error) {
	return r.client.Attendance.
		Query().
		Where(attendance.ID(id)).
		Where(attendance.DeletedAtIsNil()).
		WithEmployee().
		First(ctx)
}

// GetByEmployeeAndDate retrieves attendance by employee ID and date
func (r *AttendanceRepositoryImpl) GetByEmployeeAndDate(ctx context.Context, employeeID uint64, date time.Time) (*ent.Attendance, error) {
	// Normalize date to start of day
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	return r.client.Attendance.
		Query().
		Where(attendance.EmployeeID(employeeID)).
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		WithEmployee().
		First(ctx)
}

// Update updates an attendance record
func (r *AttendanceRepositoryImpl) Update(ctx context.Context, id uint64, req *dto.UpdateAttendanceRequest) (*ent.Attendance, error) {
	query := r.client.Attendance.UpdateOneID(id)

	if !req.CheckInTime.IsZero() {
		query = query.SetCheckInTime(req.CheckInTime)
	}
	if !req.CheckOutTime.IsZero() {
		query = query.SetCheckOutTime(req.CheckOutTime)
	}
	if req.Status != "" {
		query = query.SetStatus(attendance.Status(req.Status))
	}
	if req.Notes != "" {
		query = query.SetNotes(req.Notes)
	}
	if req.MarkedByAdmin != nil {
		query = query.SetMarkedByAdmin(*req.MarkedByAdmin)
	}

	return query.Save(ctx)
}

// Delete soft deletes an attendance record
func (r *AttendanceRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.client.Attendance.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)
}

// List retrieves attendance records with pagination and filtering
func (r *AttendanceRepositoryImpl) List(ctx context.Context, params *dto.AttendanceQueryParams) ([]*ent.Attendance, int, error) {
	query := r.client.Attendance.
		Query().
		Where(attendance.DeletedAtIsNil()).
		WithEmployee()

	// Apply filters
	if params.EmployeeID > 0 {
		query = query.Where(attendance.EmployeeID(params.EmployeeID))
	}

	if !params.StartDate.IsZero() {
		query = query.Where(attendance.AttendanceDateGTE(params.StartDate))
	}

	if !params.EndDate.IsZero() {
		query = query.Where(attendance.AttendanceDateLTE(params.EndDate))
	}

	if params.Status != "" {
		query = query.Where(attendance.StatusEQ(attendance.Status(params.Status)))
	}

	if params.IncludeWeekend != nil && !*params.IncludeWeekend {
		query = query.Where(attendance.IsWeekendEQ(false))
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count attendance records: %w", err)
	}

	// Apply pagination
	if params.Page > 0 && params.Limit > 0 {
		offset := (params.Page - 1) * params.Limit
		query = query.Offset(offset).Limit(params.Limit)
	}

	// Order by attendance_date desc, then by employee_id
	query = query.Order(ent.Desc(attendance.FieldAttendanceDate), ent.Asc(attendance.FieldEmployeeID))

	attendances, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch attendance records: %w", err)
	}

	return attendances, total, nil
}

// GetTodayAttendance retrieves today's attendance records
func (r *AttendanceRepositoryImpl) GetTodayAttendance(ctx context.Context) ([]*ent.Attendance, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	return r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		WithEmployee().
		Order(ent.Asc(attendance.FieldEmployeeID)).
		All(ctx)
}

// GetDailyAttendanceSummary calculates daily attendance summary
func (r *AttendanceRepositoryImpl) GetDailyAttendanceSummary(ctx context.Context, date time.Time) (*dto.DailyAttendanceSummary, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	// Get total active employees
	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.IsActiveEQ(true)).
		Where(employee.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count total employees: %w", err)
	}

	// Get attendance counts by status
	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attendance records: %w", err)
	}

	summary := &dto.DailyAttendanceSummary{
		Date:           startOfDay,
		TotalEmployees: totalEmployees,
	}

	// Count by status
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent:
			summary.PresentCount++
		case attendance.StatusAbsent:
			summary.AbsentCount++
		case attendance.StatusLate:
			summary.LateCount++
		case attendance.StatusHalfDay:
			summary.HalfDayCount++
		}
	}

	// Calculate percentages
	if totalEmployees > 0 {
		summary.PresentPercent = float64(summary.PresentCount) / float64(totalEmployees) * 100
		summary.AbsentPercent = float64(summary.AbsentCount) / float64(totalEmployees) * 100
	}

	return summary, nil
}

// BulkMarkAttendance marks attendance for multiple employees
func (r *AttendanceRepositoryImpl) BulkMarkAttendance(ctx context.Context, req *dto.BulkMarkAttendanceRequest) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	for _, item := range req.Attendances {
		markReq := &dto.MarkAttendanceRequest{
			EmployeeID:     item.EmployeeID,
			AttendanceDate: req.AttendanceDate,
			CheckInTime:    item.CheckInTime,
			CheckOutTime:   item.CheckOutTime,
			Status:         item.Status,
			Notes:          item.Notes,
			MarkedByAdmin:  true,
		}

		_, err := r.MarkAttendance(ctx, markReq)
		if err != nil {
			return fmt.Errorf("failed to mark attendance for employee %d: %w", item.EmployeeID, err)
		}
	}

	return tx.Commit()
}

// GetAttendanceByDateRange retrieves attendance records for an employee within a date range
func (r *AttendanceRepositoryImpl) GetAttendanceByDateRange(ctx context.Context, employeeID uint64, startDate, endDate time.Time) ([]*ent.Attendance, error) {
	return r.client.Attendance.
		Query().
		Where(attendance.EmployeeID(employeeID)).
		Where(attendance.AttendanceDateGTE(startDate)).
		Where(attendance.AttendanceDateLTE(endDate)).
		Where(attendance.DeletedAtIsNil()).
		WithEmployee().
		Order(ent.Asc(attendance.FieldAttendanceDate)).
		All(ctx)
}

// isWeekendDay checks if the given date is a weekend (Saturday or Sunday)
func isWeekendDay(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
