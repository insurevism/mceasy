package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterSalaryRoutes registers all salary calculation routes
func RegisterSalaryRoutes(e *echo.Group, controller *SalaryController) {
	// Basic CRUD operations
	e.POST("/salary/calculate", controller.CalculateSalary)
	e.GET("/salary", controller.ListSalaryCalculations)
	e.GET("/salary/:id", controller.GetSalaryCalculation)
	e.PUT("/salary/:id", controller.UpdateSalaryCalculation)
	e.DELETE("/salary/:id", controller.DeleteSalaryCalculation)

	// Bulk operations
	e.POST("/salary/calculate/bulk", controller.BulkCalculateSalary)

	// Employee-specific salary operations
	e.GET("/salary/employee/:employee_id", controller.GetSalaryCalculationByEmployeeAndMonth)

	// Summary operations
	e.GET("/salary/summary/monthly", controller.GetMonthlySalarySummary)
	e.GET("/salary/summary/employee/:employee_id", controller.GetEmployeeSalarySummary)
}
