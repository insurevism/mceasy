package rest

import (
	"hokusai/configs/rabbitmq/connection"
	"hokusai/ent"
	"hokusai/internal/applications/active_transaction"
	activeTransactionController "hokusai/internal/applications/active_transaction/controller"
	"hokusai/internal/applications/auth"
	authController "hokusai/internal/applications/auth/controller"
	"hokusai/internal/applications/tick_v2"
	tickV2Controller "hokusai/internal/applications/tick_v2/controller"

	"hokusai/internal/applications/daily_report"
	dailyReportController "hokusai/internal/applications/daily_report/controller"

	exampleRabbit "hokusai/internal/applications/example_rabbitmq/controller"
	"hokusai/internal/applications/health"
	"hokusai/internal/applications/health/controller"
	quotes "hokusai/internal/applications/quotes"
	quotesController "hokusai/internal/applications/quotes/controller"
	"hokusai/internal/applications/system_parameter"
	systemParameterController "hokusai/internal/applications/system_parameter/controller"
	"hokusai/internal/applications/tick"
	tickController "hokusai/internal/applications/tick/controller"
	"hokusai/internal/applications/user"
	userController "hokusai/internal/applications/user/controller"
	"hokusai/internal/applications/user_account"
	userAccountController "hokusai/internal/applications/user_account/controller"

	"hokusai/internal/applications/account_config"
	accountConfigController "hokusai/internal/applications/account_config/controller"
	"hokusai/internal/applications/news"
	newsController "hokusai/internal/applications/news/controller"
	"hokusai/internal/component/rabbitmq/producer"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client, redisClient *redis.Client, connRabbitMQ *connection.RabbitMQConnection) {

	appName := viper.GetString("application.name")

	// Swagger OpenAPI Docs
	e.GET(appName+"/swagger/*", echoSwagger.WrapHandler)

	//manual injection
	helloWorldsService := health.InitializeHealthService(connDb, redisClient)
	controller.
		NewHealthController(helloWorldsService).
		AddRoutes(e, appName)

	//injection using code gen - google wire
	systemParameterService := system_parameter.InitializedSystemParameterService(connDb, redisClient)
	systemParameterController.
		NewSystemParameterController(systemParameterService).
		AddRoutes(e, appName)

	quotesService := quotes.InitializedQuotesService()
	quotesController.
		NewQuotesController(quotesService).
		AddRoutes(e, appName)

	//add example controller:
	producerService := producer.InitializedProducer(connRabbitMQ)
	exampleRabbit.NewExampleRabbitMQController(producerService).AddRoutes(e, appName)

	//initialized credentials:

	// Auth service
	authService := auth.InitializedAuthService(connDb, redisClient)
	authController.NewAuthController(authService).AddRoutes(e, appName)

	// User service
	userService := user.InitializedUserService(connDb, redisClient)
	userController.NewUserController(userService, authService).AddRoutes(e, appName)

	// User Account service
	userAccountService := user_account.InitializedUserAccountService(connDb, redisClient)
	userAccountController.NewUserAccountController(userAccountService, authService).AddRoutes(e, appName)

	// Account Config service
	accountConfigService := account_config.InitializedAccountConfigService(connDb, redisClient)
	accountConfigController.NewAccountConfigController(accountConfigService, authService).AddRoutes(e, appName)

	// Tick Service
	tickService := tick.InitializedTickService(connDb, connRabbitMQ, redisClient)
	tickController.NewTickController(tickService, authService, producerService, redisClient).AddRoutes(e, appName)

	// Tick V2 Service
	tickV2Service := tick_v2.InitializedTickV2Service(connDb, connRabbitMQ, redisClient)
	tickV2Controller.NewTickV2Controller(tickV2Service, authService, producerService, redisClient).AddRoutes(e, appName)

	// Active Transaction service
	activeTransactionService := active_transaction.InitializedActiveTransactionService(connDb, redisClient)
	activeTransactionController.NewActiveTransactionController(activeTransactionService, authService).AddRoutes(e, appName)

	// Daily Report
	dailyReportService := daily_report.InitializedDailyReportService(connDb, redisClient)
	dailyReportController.NewDailyReportController(dailyReportService, authService).AddRoutes(e, appName)

	// News
	newsService := news.InitializedNewsService(connDb, redisClient)
	newsController.NewNewsController(newsService, authService).AddRoutes(e, appName)

}
