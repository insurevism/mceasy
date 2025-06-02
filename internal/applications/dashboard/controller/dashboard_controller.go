package controller

import (
	"net/http"
	"strconv"
	"time"

	"mceasy/internal/applications/dashboard/dto"
	"mceasy/internal/applications/dashboard/service"

	"github.com/labstack/echo/v4"
)

// DashboardController handles HTTP requests for dashboard operations
type DashboardController struct {
	dashboardService service.DashboardService
}

// NewDashboardController creates a new dashboard controller instance
func NewDashboardController(dashboardService service.DashboardService) *DashboardController {
	return &DashboardController{
		dashboardService: dashboardService,
	}
}

// GetTodayAttendanceSummary retrieves today's attendance summary
// @Summary Get today's attendance summary
// @Description Get today's attendance summary showing present, absent, and late counts
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} dto.TodayAttendanceSummary
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/today [get]
func (c *DashboardController) GetTodayAttendanceSummary(ctx echo.Context) error {
	summary, err := c.dashboardService.GetTodayAttendanceSummary(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, summary)
}

// GetDashboardOverview retrieves complete dashboard overview
// @Summary Get dashboard overview
// @Description Get complete dashboard overview with today's attendance, trends, and statistics
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} dto.DashboardOverview
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/overview [get]
func (c *DashboardController) GetDashboardOverview(ctx echo.Context) error {
	overview, err := c.dashboardService.GetDashboardOverview(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, overview)
}

// GetAttendanceTrends retrieves attendance trends
// @Summary Get attendance trends
// @Description Get attendance trends over a specified period with optional date range and period type
// @Tags dashboard
// @Accept json
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)" example(2024-01-01)
// @Param end_date query string false "End date (YYYY-MM-DD)" example(2024-01-31)
// @Param period_type query string false "Period type" Enums(daily, weekly, monthly) default(daily)
// @Success 200 {object} dto.AttendanceTrendData
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/attendance/trends [get]
func (c *DashboardController) GetAttendanceTrends(ctx echo.Context) error {
	var params dto.DashboardQueryParams

	// Parse start_date
	if startDateStr := ctx.QueryParam("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid start_date format. Use YYYY-MM-DD",
			})
		}
		params.StartDate = startDate
	}

	// Parse end_date
	if endDateStr := ctx.QueryParam("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid end_date format. Use YYYY-MM-DD",
			})
		}
		params.EndDate = endDate
	}

	// Parse period_type
	params.PeriodType = ctx.QueryParam("period_type")

	trends, err := c.dashboardService.GetAttendanceTrends(ctx.Request().Context(), params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, trends)
}

// GetSalaryTrends retrieves salary trends
// @Summary Get salary trends
// @Description Get salary trends over a specified period with optional date range
// @Tags dashboard
// @Accept json
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)" example(2024-01-01)
// @Param end_date query string false "End date (YYYY-MM-DD)" example(2024-12-31)
// @Success 200 {object} dto.SalaryTrendData
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/salary/trends [get]
func (c *DashboardController) GetSalaryTrends(ctx echo.Context) error {
	var params dto.DashboardQueryParams

	// Parse start_date
	if startDateStr := ctx.QueryParam("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid start_date format. Use YYYY-MM-DD",
			})
		}
		params.StartDate = startDate
	}

	// Parse end_date
	if endDateStr := ctx.QueryParam("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid end_date format. Use YYYY-MM-DD",
			})
		}
		params.EndDate = endDate
	}

	trends, err := c.dashboardService.GetSalaryTrends(ctx.Request().Context(), params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, trends)
}

// GetAttendanceAlerts retrieves attendance alerts
// @Summary Get attendance alerts
// @Description Get attendance alerts including late employees, absent employees, and notifications
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} dto.AttendanceAlerts
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/alerts [get]
func (c *DashboardController) GetAttendanceAlerts(ctx echo.Context) error {
	alerts, err := c.dashboardService.GetAttendanceAlerts(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, alerts)
}

// GetWeeklySummary retrieves weekly attendance summary
// @Summary Get weekly attendance summary
// @Description Get weekly attendance summary starting from a specific date (will adjust to Monday)
// @Tags dashboard
// @Accept json
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)" example(2024-01-01)
// @Success 200 {object} dto.WeeklySummary
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/weekly [get]
func (c *DashboardController) GetWeeklySummary(ctx echo.Context) error {
	startDate := time.Now()

	// Parse start_date if provided
	if startDateStr := ctx.QueryParam("start_date"); startDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid start_date format. Use YYYY-MM-DD",
			})
		}
		startDate = parsedDate
	}

	summary, err := c.dashboardService.GetWeeklySummary(ctx.Request().Context(), startDate)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, summary)
}

// GetMonthlyStats retrieves monthly statistics
// @Summary Get monthly statistics
// @Description Get monthly attendance and salary statistics for a specific month
// @Tags dashboard
// @Accept json
// @Produce json
// @Param month query string false "Month (YYYY-MM)" example(2024-01)
// @Success 200 {object} dto.MonthlyStatsSummary
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/monthly [get]
func (c *DashboardController) GetMonthlyStats(ctx echo.Context) error {
	month := time.Now()

	// Parse month if provided
	if monthStr := ctx.QueryParam("month"); monthStr != "" {
		parsedMonth, err := time.Parse("2006-01", monthStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid month format. Use YYYY-MM",
			})
		}
		month = parsedMonth
	}

	stats, err := c.dashboardService.GetMonthlyStats(ctx.Request().Context(), month)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, stats)
}

// GetSalaryOverview retrieves salary overview
// @Summary Get salary overview
// @Description Get salary overview for a specific month including calculations and statistics
// @Tags dashboard
// @Accept json
// @Produce json
// @Param month query string false "Month (YYYY-MM)" example(2024-01)
// @Success 200 {object} dto.SalaryOverviewData
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/salary/overview [get]
func (c *DashboardController) GetSalaryOverview(ctx echo.Context) error {
	month := time.Now()

	// Parse month if provided
	if monthStr := ctx.QueryParam("month"); monthStr != "" {
		parsedMonth, err := time.Parse("2006-01", monthStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid month format. Use YYYY-MM",
			})
		}
		month = parsedMonth
	}

	overview, err := c.dashboardService.GetSalaryOverview(ctx.Request().Context(), month)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, overview)
}

// GetSystemStats retrieves system statistics
// @Summary Get system statistics
// @Description Get overall system statistics including total employees, records, and uptime
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} dto.SystemStatsData
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/system [get]
func (c *DashboardController) GetSystemStats(ctx echo.Context) error {
	stats, err := c.dashboardService.GetSystemStats(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, stats)
}

// GetAttendanceTrendsByEmployee retrieves attendance trends for a specific employee
// @Summary Get employee attendance trends
// @Description Get attendance trends for a specific employee over a specified period
// @Tags dashboard
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Param start_date query string false "Start date (YYYY-MM-DD)" example(2024-01-01)
// @Param end_date query string false "End date (YYYY-MM-DD)" example(2024-01-31)
// @Param period_type query string false "Period type" Enums(daily, weekly, monthly) default(daily)
// @Success 200 {object} dto.AttendanceTrendData
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /dashboard/attendance/trends/employee/{employee_id} [get]
func (c *DashboardController) GetAttendanceTrendsByEmployee(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid employee ID",
		})
	}

	var params dto.DashboardQueryParams
	params.EmployeeID = employeeID

	// Parse start_date
	if startDateStr := ctx.QueryParam("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid start_date format. Use YYYY-MM-DD",
			})
		}
		params.StartDate = startDate
	}

	// Parse end_date
	if endDateStr := ctx.QueryParam("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid end_date format. Use YYYY-MM-DD",
			})
		}
		params.EndDate = endDate
	}

	// Parse period_type
	params.PeriodType = ctx.QueryParam("period_type")

	trends, err := c.dashboardService.GetAttendanceTrends(ctx.Request().Context(), params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, trends)
}
