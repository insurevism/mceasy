package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"

	_ "mceasy/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

func main() {
	loadConfig()

	flags.Usage = usage
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	args := flags.Args()
	if len(args) < 2 {
		flags.Usage()
		return
	}

	driver := args[0]
	command := args[1]

	// Ensure MySQL is supported
	if driver != "mysql" {
		log.Fatalf("%q driver not supported. Use 'mysql'.", driver)
	}

	// Get database connection string from environment variables
	dbSource := getMySQLDSN()
	if dbSource == "" {
		log.Fatal("MySQL DSN could not be constructed from environment variables")
	}

	// Open MySQL database connection
	db, err := sql.Open(driver, dbSource)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Run Goose migration command
	executeCommand(args, command, db)
}

// loadConfig initializes Viper and reads secret.env
func loadConfig() {
	viper.SetConfigFile("secret.env") // Set config file name
	viper.SetConfigType("env")        // Define file type as .env
	viper.AutomaticEnv()              // Enable environment variable support

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

// getMySQLDSN constructs the MySQL DSN from environment variables
func getMySQLDSN() string {
	host := viper.GetString("db.configs.host")
	port := viper.GetString("db.configs.port")
	user := viper.GetString("db.configs.username")
	password := viper.GetString("db.configs.password")
	dbname := viper.GetString("db.configs.database")

	if host == "" || port == "" || user == "" || dbname == "" {
		log.Fatal("Missing required database environment variables")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname)
}

// executeCommand runs the Goose command
func executeCommand(args []string, command string, db *sql.DB) {
	var arguments []string
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.RunWithOptions(command, db, *dir, arguments, goose.WithAllowMissing()); err != nil {
		log.Fatalf("Goose run error: %v", err)
	}
}

// usage prints help message
func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: go run main.go DRIVER COMMAND
Drivers:
    mysql
Examples:
    go run main.go mysql up
    go run main.go mysql down
    go run main.go mysql status
    go run main.go mysql create migration_name sql
`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    down                 Roll back the latest migration
    status               Show migration status
    create NAME TYPE     Create a new migration file (TYPE: sql or go)
`
)
