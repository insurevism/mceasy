package rest

import (
	"mceasy/configs/rabbitmq/connection"
	"mceasy/ent"
	"mceasy/internal/applications/active_transaction"
	activeTransactionController "mceasy/internal/applications/active_transaction/controller"
	"mceasy/internal/applications/auth"
	authController "mceasy/internal/applications/auth/controller"
	"mceasy/internal/applications/tick_v2"
	tickV2Controller "mceasy/internal/applications/tick_v2/controller"

	"mceasy/internal/applications/daily_report"
	dailyReportController "mceasy/internal/applications/daily_report/controller"

	exampleRabbit "mceasy/internal/applications/example_rabbitmq/controller"
	"mceasy/internal/applications/health"
	"mceasy/internal/applications/health/controller"
	quotes "mceasy/internal/applications/quotes"
	quotesController "mceasy/internal/applications/quotes/controller"
	"mceasy/internal/applications/system_parameter"
	systemParameterController "mceasy/internal/applications/system_parameter/controller"
	"mceasy/internal/applications/tick"
	tickController "mceasy/internal/applications/tick/controller"
	"mceasy/internal/applications/user"
	userController "mceasy/internal/applications/user/controller"
	"mceasy/internal/applications/user_account"
	userAccountController "mceasy/internal/applications/user_account/controller"

	"mceasy/internal/applications/account_config"
	accountConfigController "mceasy/internal/applications/account_config/controller"
	"mceasy/internal/applications/news"
	newsController "mceasy/internal/applications/news/controller"
	"mceasy/internal/component/rabbitmq/producer"

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
