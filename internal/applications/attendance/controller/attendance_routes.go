package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterAttendanceRoutes registers all attendance routes
func RegisterAttendanceRoutes(e *echo.Group, controller *AttendanceController) {
	// Basic CRUD operations
	e.POST("/attendance", controller.MarkAttendance)
	e.GET("/attendance", controller.ListAttendance)
	e.GET("/attendance/:id", controller.GetAttendance)
	e.PUT("/attendance/:id", controller.UpdateAttendance)
	e.DELETE("/attendance/:id", controller.DeleteAttendance)

	// Special attendance operations
	e.GET("/attendance/today", controller.GetTodayAttendance)
	e.GET("/attendance/summary", controller.GetDailyAttendanceSummary)
	e.POST("/attendance/bulk", controller.BulkMarkAttendance)

	// Employee-specific attendance operations
	e.POST("/attendance/checkin/:employee_id", controller.CheckInEmployee)
	e.POST("/attendance/checkout/:employee_id", controller.CheckOutEmployee)
	e.GET("/attendance/employee/:employee_id", controller.GetEmployeeAttendanceByDate)
	e.GET("/attendance/employee/:employee_id/history", controller.GetEmployeeAttendanceHistory)
}
