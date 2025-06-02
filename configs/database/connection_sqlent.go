package database

import (
	"database/sql"
	"fmt"
	"time"

	"hokusai/configs/credential"
	"hokusai/ent"

	entSql "entgo.io/ent/dialect/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func NewSqlEntClient() *ent.Client {
	// Adjust the DSN format for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		credential.GetString("db.configs.host"),
		credential.GetString("db.configs.port"),
		credential.GetString("db.configs.username"),
		credential.GetString("db.configs.password"),
		credential.GetString("db.configs.database"))

	log.Infof("DSN=", dsn) // For debugging only, ensure credentials are not logged in production

	// Open a PostgreSQL connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to DB: %v", err)
	}

	// Set database connection settings
	db.SetMaxIdleConns(credential.GetInt("db.configs.maxIdleConn"))
	db.SetMaxOpenConns(credential.GetInt("db.configs.maxOpenConn"))
	db.SetConnMaxLifetime(time.Hour)

	// Initialize the Ent client with PostgreSQL driver
	drv := entSql.OpenDB("postgres", db)
	client := ent.NewClient(ent.Driver(drv))

	// Ensure the client is properly initialized
	if client == nil {
		log.Fatalf("failed to initialize Ent client")
	}

	// Determine app mode
	appMode := credential.GetString("application.mode")
	if appMode == "prod" {
		log.Info("initialized database connection: PRODUCTION")
	} else {
		log.Info("initialized database connection: DEVELOPMENT")
		client = ent.NewClient(ent.Driver(drv), ent.Debug()) // Enable debug in non-prod mode
	}

	log.Info("Database initialization successful")

	// Setup hooks if needed
	SetupHooks(client)

	return client
}
