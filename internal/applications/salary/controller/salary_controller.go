package controller

import (
	"net/http"
	"strconv"
	"time"

	"mceasy/internal/applications/salary/dto"
	"mceasy/internal/applications/salary/service"

	"github.com/labstack/echo/v4"
)

// SalaryController handles HTTP requests for salary calculation operations
type SalaryController struct {
	salaryService service.SalaryService
}

// NewSalaryController creates a new salary controller instance
func NewSalaryController(salaryService service.SalaryService) *SalaryController {
	return &SalaryController{
		salaryService: salaryService,
	}
}

// CalculateSalary calculates monthly salary for an employee
// @Summary Calculate monthly salary
// @Description Calculate monthly salary for an employee based on attendance
// @Tags salary
// @Accept json
// @Produce json
// @Param salary body dto.CalculateSalaryRequest true "Salary calculation data"
// @Success 201 {object} dto.SalaryCalculationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/calculate [post]
func (c *SalaryController) CalculateSalary(ctx echo.Context) error {
	var req dto.CalculateSalaryRequest
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

	calculation, err := c.salaryService.CalculateSalary(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to calculate salary",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, calculation)
}

// GetSalaryCalculation retrieves a salary calculation by ID
// @Summary Get salary calculation by ID
// @Description Get salary calculation details by ID
// @Tags salary
// @Accept json
// @Produce json
// @Param id path int true "Salary Calculation ID"
// @Success 200 {object} dto.SalaryCalculationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /salary/{id} [get]
func (c *SalaryController) GetSalaryCalculation(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid salary calculation ID",
			"message": "Salary calculation ID must be a valid number",
		})
	}

	calculation, err := c.salaryService.GetSalaryCalculationByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Salary calculation not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, calculation)
}

// UpdateSalaryCalculation updates a salary calculation
// @Summary Update salary calculation
// @Description Update salary calculation details
// @Tags salary
// @Accept json
// @Produce json
// @Param id path int true "Salary Calculation ID"
// @Param salary body dto.UpdateSalaryCalculationRequest true "Salary calculation data"
// @Success 200 {object} dto.SalaryCalculationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/{id} [put]
func (c *SalaryController) UpdateSalaryCalculation(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid salary calculation ID",
			"message": "Salary calculation ID must be a valid number",
		})
	}

	var req dto.UpdateSalaryCalculationRequest
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

	calculation, err := c.salaryService.UpdateSalaryCalculation(ctx.Request().Context(), id, &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to update salary calculation",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, calculation)
}

// DeleteSalaryCalculation deletes a salary calculation
// @Summary Delete salary calculation
// @Description Soft delete a salary calculation
// @Tags salary
// @Accept json
// @Produce json
// @Param id path int true "Salary Calculation ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/{id} [delete]
func (c *SalaryController) DeleteSalaryCalculation(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid salary calculation ID",
			"message": "Salary calculation ID must be a valid number",
		})
	}

	err = c.salaryService.DeleteSalaryCalculation(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to delete salary calculation",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Salary calculation deleted successfully",
	})
}

// ListSalaryCalculations retrieves salary calculations with pagination and filtering
// @Summary List salary calculations
// @Description Get list of salary calculations with pagination and filtering
// @Tags salary
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param employee_id query int false "Employee ID filter"
// @Param calculation_month query string false "Calculation month filter (YYYY-MM-DD)"
// @Param start_month query string false "Start month filter (YYYY-MM-DD)"
// @Param end_month query string false "End month filter (YYYY-MM-DD)"
// @Success 200 {object} dto.SalaryCalculationListResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary [get]
func (c *SalaryController) ListSalaryCalculations(ctx echo.Context) error {
	var params dto.SalaryCalculationQueryParams
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

	calculations, err := c.salaryService.ListSalaryCalculations(ctx.Request().Context(), &params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to list salary calculations",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, calculations)
}

// GetMonthlySalarySummary retrieves monthly salary summary
// @Summary Get monthly salary summary
// @Description Get salary summary for a specific month
// @Tags salary
// @Accept json
// @Produce json
// @Param month query string true "Month (YYYY-MM-DD)"
// @Success 200 {object} dto.MonthlySalarySummary
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/summary/monthly [get]
func (c *SalaryController) GetMonthlySalarySummary(ctx echo.Context) error {
	monthStr := ctx.QueryParam("month")
	if monthStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Missing required parameter",
			"message": "month parameter is required",
		})
	}

	month, err := time.Parse("2006-01-02", monthStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid month format",
			"message": "Month must be in YYYY-MM-DD format",
		})
	}

	summary, err := c.salaryService.GetMonthlySalarySummary(ctx.Request().Context(), month)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get monthly salary summary",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, summary)
}

// GetEmployeeSalarySummary retrieves salary summary for an employee across months
// @Summary Get employee salary summary
// @Description Get salary summary for a specific employee within date range
// @Tags salary
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Param start_month query string true "Start month (YYYY-MM-DD)"
// @Param end_month query string true "End month (YYYY-MM-DD)"
// @Success 200 {object} dto.EmployeeSalarySummary
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/summary/employee/{employee_id} [get]
func (c *SalaryController) GetEmployeeSalarySummary(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	startMonthStr := ctx.QueryParam("start_month")
	endMonthStr := ctx.QueryParam("end_month")

	if startMonthStr == "" || endMonthStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Missing required parameters",
			"message": "start_month and end_month are required",
		})
	}

	startMonth, err := time.Parse("2006-01-02", startMonthStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid start_month format",
			"message": "Month must be in YYYY-MM-DD format",
		})
	}

	endMonth, err := time.Parse("2006-01-02", endMonthStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid end_month format",
			"message": "Month must be in YYYY-MM-DD format",
		})
	}

	summary, err := c.salaryService.GetEmployeeSalarySummary(ctx.Request().Context(), employeeID, startMonth, endMonth)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to get employee salary summary",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, summary)
}

// BulkCalculateSalary calculates salary for multiple employees
// @Summary Bulk calculate salary
// @Description Calculate salary for multiple employees at once
// @Tags salary
// @Accept json
// @Produce json
// @Param salary body dto.BulkCalculateSalaryRequest true "Bulk salary calculation data"
// @Success 200 {array} dto.SalaryCalculationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /salary/calculate/bulk [post]
func (c *SalaryController) BulkCalculateSalary(ctx echo.Context) error {
	var req dto.BulkCalculateSalaryRequest
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

	calculations, err := c.salaryService.BulkCalculateSalary(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to bulk calculate salary",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, calculations)
}

// GetSalaryCalculationByEmployeeAndMonth retrieves salary calculation by employee and month
// @Summary Get salary calculation by employee and month
// @Description Get salary calculation for a specific employee and month
// @Tags salary
// @Accept json
// @Produce json
// @Param employee_id path int true "Employee ID"
// @Param month query string true "Month (YYYY-MM-DD)"
// @Success 200 {object} dto.SalaryCalculationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /salary/employee/{employee_id} [get]
func (c *SalaryController) GetSalaryCalculationByEmployeeAndMonth(ctx echo.Context) error {
	employeeIDStr := ctx.Param("employee_id")
	employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid employee ID",
			"message": "Employee ID must be a valid number",
		})
	}

	monthStr := ctx.QueryParam("month")
	if monthStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Missing required parameter",
			"message": "month parameter is required",
		})
	}

	month, err := time.Parse("2006-01-02", monthStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid month format",
			"message": "Month must be in YYYY-MM-DD format",
		})
	}

	calculation, err := c.salaryService.GetSalaryCalculationByEmployeeAndMonth(ctx.Request().Context(), employeeID, month)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "Salary calculation not found",
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, calculation)
}
