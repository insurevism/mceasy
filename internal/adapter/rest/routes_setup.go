package rest

import (
	"mceasy/ent"
	"mceasy/internal/applications/attendance"
	attendanceController "mceasy/internal/applications/attendance/controller"
	"mceasy/internal/applications/auth"
	authController "mceasy/internal/applications/auth/controller"
	"mceasy/internal/applications/employee"
	employeeController "mceasy/internal/applications/employee/controller"
	"mceasy/internal/applications/health"
	"mceasy/internal/applications/health/controller"
	"mceasy/internal/applications/salary"
	salaryController "mceasy/internal/applications/salary/controller"
	"mceasy/internal/applications/user"
	userController "mceasy/internal/applications/user/controller"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client, redisClient *redis.Client) {

	appName := viper.GetString("application.name")

	// Swagger OpenAPI Docs
	e.GET(appName+"/swagger/*", echoSwagger.WrapHandler)

	// API group
	api := e.Group(appName + "/api/v1")

	// Health service (manual injection)
	helloWorldsService := health.InitializeHealthService(connDb, redisClient)
	controller.
		NewHealthController(helloWorldsService).
		AddRoutes(e, appName)

	// Auth service
	authService := auth.InitializedAuthService(connDb, redisClient)
	authController.NewAuthController(authService).AddRoutes(e, appName)

	// User service
	userService := user.InitializedUserService(connDb, redisClient)
	userController.NewUserController(userService, authService).AddRoutes(e, appName)

	// Employee service
	employeeService := employee.InitializedEmployeeService(connDb, redisClient)
	employeeCtrl := employeeController.NewEmployeeController(employeeService)
	employeeController.RegisterEmployeeRoutes(api, employeeCtrl)

	// Attendance service
	attendanceService := attendance.InitializedAttendanceService(connDb, redisClient)
	attendanceCtrl := attendanceController.NewAttendanceController(attendanceService)
	attendanceController.RegisterAttendanceRoutes(api, attendanceCtrl)

	// Salary service
	salaryService := salary.InitializedSalaryService(connDb, redisClient)
	salaryCtrl := salaryController.NewSalaryController(salaryService)
	salaryController.RegisterSalaryRoutes(api, salaryCtrl)
}
