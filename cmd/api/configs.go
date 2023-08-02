package api

import (
	"flag"
	"fmt"
	"os"

	"github.com/hossein1376/kala/cmd"
)

func newConfig() (*cmd.Config, error) {
	var cfg cmd.Config

	// getting settings from environment variables
	env := os.Getenv("ENVIRONMENT")
	port := os.Getenv("SERVER_PORT")
	jwtSecret := os.Getenv("JWT_SECRET")

	sqlUsername := os.Getenv("SQL_USERNAME")
	sqlHost := os.Getenv("SQL_HOST")
	sqlPassword := os.Getenv("SQL_PASSWORD")
	sqlName := os.Getenv("SQL_NAME")
	sqlPort := os.Getenv("SQL_PORT")

	// getting settings from command line
	// Note that flag arguments override environment variables, if provided
	flag.StringVar(&cfg.Environment, "env", env, "Service Environment")
	flag.StringVar(&cfg.Port, "port", port, "Service Port")
	flag.StringVar(&cfg.JWTSecret, "jwt-secret", jwtSecret, "JWT Secret")

	flag.StringVar(&cfg.Sql.Port, "sql-port", sqlPort, "SQL Database Port")
	flag.StringVar(&cfg.Sql.Username, "sql-username", sqlUsername, "SQL Database Username")
	flag.StringVar(&cfg.Sql.Host, "sql-host", sqlHost, "SQL Database Host")
	flag.StringVar(&cfg.Sql.Password, "sql-password", sqlPassword, "SQL Database Password")
	flag.StringVar(&cfg.Sql.Name, "sql-name", sqlName, "SQL Database Name")

	flag.Parse()

	err := verifyConfigs(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func verifyConfigs(cfg *cmd.Config) error {
	if cfg.Port == "" {
		return fmt.Errorf("port must be specified")
	}
	if cfg.Environment == "" {
		return fmt.Errorf("environment must be specified")
	}
	if cfg.JWTSecret == "" {
		return fmt.Errorf("JWT secret must be specified")
	}

	if cfg.Sql.Host == "" {
		return fmt.Errorf("sql host must be specified")
	}
	if cfg.Sql.Port == "" {
		return fmt.Errorf("sql port must be specified")
	}
	if cfg.Sql.Name == "" {
		return fmt.Errorf("sql name must be specified")
	}
	if cfg.Sql.Username == "" {
		return fmt.Errorf("sql username must be specified")
	}

	return nil
}
