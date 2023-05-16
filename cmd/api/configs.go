package api

import (
	"flag"
	"os"

	"kala/internal/structure"
)

func newConfig() structure.Config {
	var cfg structure.Config

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
	flag.StringVar(&cfg.Environment, "env", env, "Service Environment")
	flag.StringVar(&cfg.Port, "port", port, "Service Port")

	flag.StringVar(&cfg.Sql.Port, "sql-port", sqlPort, "SQL Database Port")
	flag.StringVar(&cfg.Sql.Username, "sql-username", sqlUsername, "SQL Database Username")
	flag.StringVar(&cfg.Sql.Host, "sql-host", sqlHost, "SQL Database Host")
	flag.StringVar(&cfg.Sql.Password, "sql-password", sqlPassword, "SQL Database Password")
	flag.StringVar(&cfg.Sql.Name, "sql-name", sqlName, "SQL Database Name")

	flag.Parse()
	return cfg
}
