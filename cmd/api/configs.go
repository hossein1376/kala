package api

import (
	"flag"
	"os"
)

type config struct {
	environment string
	port        string
	sql         sqlDB
}

func newConfig() config {
	var cfg config

	// getting settings from environment variables
	env := os.Getenv("ENVIRONMENT")
	port := os.Getenv("SERVER_PORT")

	sqlUsername := os.Getenv("SQL_USERNAME")
	sqlHost := os.Getenv("SQL_HOST")
	sqlPassword := os.Getenv("SQL_PASSWORD")
	sqlName := os.Getenv("SQL_NAME")
	sqlPort := os.Getenv("SQL_PORT")

	// getting settings from command line
	// Note that flag arguments override environment variables, if provided
	flag.StringVar(&cfg.environment, "env", env, "Service Environment")
	flag.StringVar(&cfg.port, "port", port, "Service Port")

	flag.StringVar(&cfg.sql.port, "sql-port", sqlPort, "SQL Database Port")
	flag.StringVar(&cfg.sql.username, "sql-username", sqlUsername, "SQL Database Username")
	flag.StringVar(&cfg.sql.host, "sql-host", sqlHost, "SQL Database Host")
	flag.StringVar(&cfg.sql.password, "sql-password", sqlPassword, "SQL Database Password")
	flag.StringVar(&cfg.sql.name, "sql-name", sqlName, "SQL Database Name")

	flag.Parse()
	return cfg
}
