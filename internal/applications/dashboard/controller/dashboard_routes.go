package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterDashboardRoutes registers all dashboard routes
func RegisterDashboardRoutes(e *echo.Group, controller *DashboardController) {
	// Core dashboard endpoints
	e.GET("/dashboard/today", controller.GetTodayAttendanceSummary)
	e.GET("/dashboard/overview", controller.GetDashboardOverview)
	e.GET("/dashboard/system", controller.GetSystemStats)

	// Attendance-related dashboard endpoints
	e.GET("/dashboard/attendance/trends", controller.GetAttendanceTrends)
	e.GET("/dashboard/attendance/trends/employee/:employee_id", controller.GetAttendanceTrendsByEmployee)
	e.GET("/dashboard/alerts", controller.GetAttendanceAlerts)

	// Salary-related dashboard endpoints
	e.GET("/dashboard/salary/trends", controller.GetSalaryTrends)
	e.GET("/dashboard/salary/overview", controller.GetSalaryOverview)

	// Time-based summary endpoints
	e.GET("/dashboard/weekly", controller.GetWeeklySummary)
	e.GET("/dashboard/monthly", controller.GetMonthlyStats)
}
