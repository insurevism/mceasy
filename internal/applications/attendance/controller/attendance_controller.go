package controller

import (
	"net/http"
	"strconv"
	"time"

	"mceasy/internal/applications/attendance/dto"
	"mceasy/internal/applications/attendance/service"

	"github.com/labstack/echo/v4"
)

// AttendanceController handles HTTP requests for attendance operations
type AttendanceController struct {
	attendanceService service.AttendanceService
}

// NewAttendanceController creates a new attendance controller instance
func NewAttendanceController(attendanceService service.AttendanceService) *AttendanceController {
	return &AttendanceController{
		attendanceService: attendanceService,
	}
}

// MarkAttendance marks attendance for an employee
// @Summary Mark attendance
// @Description Mark attendance for an employee
// @Tags attendance
// @Accept json
// @Produce json
// @Param attendance body dto.MarkAttendanceRequest true "Attendance data"
// @Success 201 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance [post]
func (c *AttendanceController) MarkAttendance(ctx echo.Context) error {
	var req dto.MarkAttendanceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	attendance, err := c.attendanceService.MarkAttendance(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to mark attendance",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, attendance)
}

// GetAttendance retrieves an attendance record by ID
// @Summary Get attendance by ID
// @Description Get attendance record details by ID
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /attendance/{id} [get]
func (c *AttendanceController) GetAttendance(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid attendance ID",
			"message": "Attendance ID must be a valid number",
		})
	}

	attendance, err := c.attendanceService.GetAttendanceByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Attendance record not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// UpdateAttendance updates an attendance record
// @Summary Update attendance
// @Description Update attendance record details
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Param attendance body dto.UpdateAttendanceRequest true "Attendance data"
// @Success 200 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/{id} [put]
func (c *AttendanceController) UpdateAttendance(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid attendance ID",
			"message": "Attendance ID must be a valid number",
		})
	}

	var req dto.UpdateAttendanceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	attendance, err := c.attendanceService.UpdateAttendance(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to update attendance",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// DeleteAttendance deletes an attendance record
// @Summary Delete attendance
// @Description Soft delete an attendance record
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/{id} [delete]
func (c *AttendanceController) DeleteAttendance(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid attendance ID",
			"message": "Attendance ID must be a valid number",
		})
	}

	err = c.attendanceService.DeleteAttendance(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to delete attendance",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Attendance record deleted successfully",
	})
}

// ListAttendance retrieves attendance records with pagination and filtering
// @Summary List attendance records
// @Description Get list of attendance records with pagination and filtering
// @Tags attendance
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param employee_id query int false "Employee ID filter"
// @Param start_date query string false "Start date filter (YYYY-MM-DD)"
// @Param end_date query string false "End date filter (YYYY-MM-DD)"
// @Param status query string false "Status filter" Enums(present, absent, late, half_day)
// @Param include_weekend query bool false "Include weekend records"
// @Success 200 {object} dto.AttendanceListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance [get]
func (c *AttendanceController) ListAttendance(ctx echo.Context) error {
	var params dto.AttendanceQueryParams
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid query parameters",
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(&params); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	attendance, err := c.attendanceService.ListAttendance(ctx.Request().Context(), &params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to list attendance records",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// GetTodayAttendance retrieves today's attendance records
// @Summary Get today's attendance
// @Description Get list of today's attendance records
// @Tags attendance
// @Accept json
// @Produce json
// @Success 200 {array} dto.AttendanceResponse
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/today [get]
func (c *AttendanceController) GetTodayAttendance(ctx echo.Context) error {
	attendance, err := c.attendanceService.GetTodayAttendance(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get today's attendance",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// GetDailyAttendanceSummary retrieves daily attendance summary
// @Summary Get daily attendance summary
// @Description Get attendance summary for a specific date
// @Tags attendance
// @Accept json
// @Produce json
// @Param date query string false "Date (YYYY-MM-DD), defaults to today"
// @Success 200 {object} dto.DailyAttendanceSummary
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/summary [get]
func (c *AttendanceController) GetDailyAttendanceSummary(ctx echo.Context) error {
	dateStr := ctx.QueryParam("date")
	var date time.Time
	var err error

	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error":   "Invalid date format",
				"message": "Date must be in YYYY-MM-DD format",
			})
		}
	}

	summary, err := c.attendanceService.GetDailyAttendanceSummary(ctx.Request().Context(), date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get daily attendance summary",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, summary)
}

// BulkMarkAttendance marks attendance for multiple employees
// @Summary Bulk mark attendance
// @Description Mark attendance for multiple employees at once
// @Tags attendance
// @Accept json
// @Produce json
// @Param attendance body dto.BulkMarkAttendanceRequest true "Bulk attendance data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/bulk [post]
func (c *AttendanceController) BulkMarkAttendance(ctx echo.Context) error {
	var req dto.BulkMarkAttendanceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Validation failed",
			"message": err.Error(),
		})
	}

	err := c.attendanceService.BulkMarkAttendance(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to bulk mark attendance",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Bulk attendance marked successfully",
	})
}

// CheckInEmployee marks an employee as present with check-in time
// @Summary Employee check-in
// @Description Mark employee check-in for today
// @Tags attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Success 200 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/checkin/{employee_id} [post]
func (c *AttendanceController) CheckInEmployee(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	checkInTime := time.Now()
	attendance, err := c.attendanceService.CheckInEmployee(ctx.Request().Context(), employeeID, checkInTime)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to check in employee",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// CheckOutEmployee updates check-out time for an employee
// @Summary Employee check-out
// @Description Mark employee check-out for today
// @Tags attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Success 200 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/checkout/{employee_id} [post]
func (c *AttendanceController) CheckOutEmployee(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	checkOutTime := time.Now()
	attendance, err := c.attendanceService.CheckOutEmployee(ctx.Request().Context(), employeeID, checkOutTime)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to check out employee",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// GetEmployeeAttendanceHistory retrieves attendance history for an employee
// @Summary Get employee attendance history
// @Description Get attendance history for a specific employee within date range
// @Tags attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /attendance/employee/{employee_id}/history [get]
func (c *AttendanceController) GetEmployeeAttendanceHistory(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	startDateStr := ctx.QueryParam("start_date")
	endDateStr := ctx.QueryParam("end_date")

	if startDateStr == "" || endDateStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Missing required parameters",
			"message": "start_date and end_date are required",
		})
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid start_date format",
			"message": "Date must be in YYYY-MM-DD format",
		})
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid end_date format",
			"message": "Date must be in YYYY-MM-DD format",
		})
	}

	attendance, err := c.attendanceService.GetEmployeeAttendanceHistory(ctx.Request().Context(), employeeID, startDate, endDate)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get employee attendance history",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}

// GetEmployeeAttendanceByDate retrieves attendance for an employee on a specific date
// @Summary Get employee attendance by date
// @Description Get attendance record for a specific employee and date
// @Tags attendance
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {object} dto.AttendanceResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /attendance/employee/{employee_id} [get]
func (c *AttendanceController) GetEmployeeAttendanceByDate(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	dateStr := ctx.QueryParam("date")
	if dateStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Missing required parameter",
			"message": "date parameter is required",
		})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid date format",
			"message": "Date must be in YYYY-MM-DD format",
		})
	}

	attendance, err := c.attendanceService.GetAttendanceByEmployeeAndDate(ctx.Request().Context(), employeeID, date)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Attendance record not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, attendance)
}
