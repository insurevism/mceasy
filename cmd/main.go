package main

import (
	"context"
	"errors"
	"fmt"
	"mceasy/configs"
	"mceasy/configs/cache"
	"mceasy/configs/credential"
	"mceasy/configs/database"
	"mceasy/configs/swagger"
	"mceasy/configs/validator"
	"mceasy/ent"
	restApi "mceasy/internal/adapter/rest"
	"mceasy/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	_ "mceasy/cmd/docs"
)

//	@title			Attendance Management System
//	@version		1.0.0
//	@description	Attendance Management System for Employee Tracking

//	@contact.url	https://example.com

// @host		localhost:8888
// @basePath	/attendance-system
func main() {
	e := echo.New()
	fmt.Println("initialized echo framework")

	configs.InitGeneralEnv(e)
	credential.InitCredentialEnv(e)
	configs.SetupLogger(e)
	configs.SetupZeroLogger()

	middleware.SetupMiddlewares(e)

	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	dbConnection := database.NewSqlEntClient()

	//from docs define close on this function, but will impact cant create DB session on repository:
	defer func(dbConnection *ent.Client) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatalf("error initialized database configuration=%v", err)
		}
	}(dbConnection)

	//configuration for redis client:
	redisConnection := cache.NewRedisClient()

	//configuration for redis client, for close connection:
	defer func() {
		err := redisConnection.Close()
		if err != nil {
			log.Fatalf("Error closing Redis connection: %v", err)
		}
	}()

	//setup swagger:
	swagger.InitSwagger()

	//setup router
	restApi.SetupRouteHandler(e, dbConnection, redisConnection)

	port := viper.GetString("application.port")

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
			e.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
