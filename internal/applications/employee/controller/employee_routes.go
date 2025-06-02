package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterEmployeeRoutes registers all employee routes
func RegisterEmployeeRoutes(e *echo.Group, controller *EmployeeController) {
	// Employee management routes
	e.POST("/employees", controller.CreateEmployee)
	e.GET("/employees", controller.ListEmployees)
	e.GET("/employees/active", controller.GetActiveEmployees)
	e.GET("/employees/:id", controller.GetEmployee)
	e.GET("/employees/by-employee-id/:employee_id", controller.GetEmployeeByEmployeeID)
	e.PUT("/employees/:id", controller.UpdateEmployee)
	e.DELETE("/employees/:id", controller.DeleteEmployee)
}
