package controller

import (
	"net/http"
	"strconv"

	"mceasy/internal/applications/employee/dto"
	"mceasy/internal/applications/employee/service"

	"github.com/labstack/echo/v4"
)

// EmployeeController handles HTTP requests for employee operations
type EmployeeController struct {
	employeeService service.EmployeeService
}

// NewEmployeeController creates a new employee controller instance
func NewEmployeeController(employeeService service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

// CreateEmployee creates a new employee
// @Summary Create a new employee
// @Description Create a new employee with unique employee ID
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body dto.CreateEmployeeRequest true "Employee data"
// @Success 201 {object} dto.EmployeeResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees [post]
func (c *EmployeeController) CreateEmployee(ctx echo.Context) error {
	var req dto.CreateEmployeeRequest
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

	employee, err := c.employeeService.CreateEmployee(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to create employee",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, employee)
}

// GetEmployee retrieves an employee by ID
// @Summary Get employee by ID
// @Description Get employee details by ID
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} dto.EmployeeResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /employees/{id} [get]
func (c *EmployeeController) GetEmployee(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	employee, err := c.employeeService.GetEmployeeByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Employee not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, employee)
}

// GetEmployeeByEmployeeID retrieves an employee by employee ID
// @Summary Get employee by employee ID
// @Description Get employee details by employee ID (e.g., EMP-0001)
// @Tags employees
// @Accept json
// @Produce json
// @Param employee_id path string true "Employee ID (e.g., EMP-0001)"
// @Success 200 {object} dto.EmployeeResponse
// @Failure 404 {object} map[string]interface{}
// @Router /employees/by-employee-id/{employee_id} [get]
func (c *EmployeeController) GetEmployeeByEmployeeID(ctx echo.Context) error {
	employeeID := ctx.Param("employee_id")

	employee, err := c.employeeService.GetEmployeeByEmployeeID(ctx.Request().Context(), employeeID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Employee not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, employee)
}

// UpdateEmployee updates an employee
// @Summary Update employee
// @Description Update employee details
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body dto.UpdateEmployeeRequest true "Employee data"
// @Success 200 {object} dto.EmployeeResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees/{id} [put]
func (c *EmployeeController) UpdateEmployee(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	var req dto.UpdateEmployeeRequest
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

	employee, err := c.employeeService.UpdateEmployee(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to update employee",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, employee)
}

// DeleteEmployee deletes an employee
// @Summary Delete employee
// @Description Soft delete an employee
// @Tags employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees/{id} [delete]
func (c *EmployeeController) DeleteEmployee(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	err = c.employeeService.DeleteEmployee(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to delete employee",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Employee deleted successfully",
	})
}

// ListEmployees retrieves employees with pagination and filtering
// @Summary List employees
// @Description Get list of employees with pagination and filtering
// @Tags employees
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search term"
// @Param department query string false "Department filter"
// @Param is_active query bool false "Active status filter"
// @Success 200 {object} dto.EmployeeListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /employees [get]
func (c *EmployeeController) ListEmployees(ctx echo.Context) error {
	var params dto.EmployeeQueryParams
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

	employees, err := c.employeeService.ListEmployees(ctx.Request().Context(), &params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to list employees",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, employees)
}

// GetActiveEmployees retrieves all active employees
// @Summary Get active employees
// @Description Get list of all active employees
// @Tags employees
// @Accept json
// @Produce json
// @Success 200 {array} dto.EmployeeResponse
// @Failure 500 {object} map[string]interface{}
// @Router /employees/active [get]
func (c *EmployeeController) GetActiveEmployees(ctx echo.Context) error {
	employees, err := c.employeeService.GetActiveEmployees(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get active employees",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, employees)
}
