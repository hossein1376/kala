package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/hossein1376/kala/config"
)

func newConfig() (*config.Config, error) {
	var cfg config.Config

	// getting settings from environment variables
	env := os.Getenv("ENVIRONMENT")
	port := os.Getenv("SERVER_PORT")

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpire := os.Getenv("JWT_SECRET")

	logLevel := os.Getenv("LOG_LEVEL")

	sqlUsername := os.Getenv("SQL_USERNAME")
	sqlHost := os.Getenv("SQL_HOST")
	sqlPassword := os.Getenv("SQL_PASSWORD")
	sqlName := os.Getenv("SQL_NAME")
	sqlPort := os.Getenv("SQL_PORT")

	cfgFile := os.Getenv("CONFIG_FILE")

	// getting settings from command line
	// Note that flag arguments override environment variables, if provided
	flag.StringVar(&cfg.Environment, "env", env, "Service Environment: dev - prod")
	flag.StringVar(&cfg.Port, "port", port, "Service Port")

	flag.StringVar(&cfg.JWT.Secret, "jwt-secret", jwtSecret, "JWT Secret")
	flag.StringVar(&cfg.JWT.Expire, "jwt-expire", jwtExpire, "JWT Expire")

	flag.StringVar(&cfg.Logger.Level, "log-level", logLevel, "Log Level")

	flag.StringVar(&cfg.DB.Sql.Port, "sql-port", sqlPort, "SQL Database Port")
	flag.StringVar(&cfg.DB.Sql.Username, "sql-username", sqlUsername, "SQL Database Username")
	flag.StringVar(&cfg.DB.Sql.Host, "sql-host", sqlHost, "SQL Database Host")
	flag.StringVar(&cfg.DB.Sql.Password, "sql-password", sqlPassword, "SQL Database Password")
	flag.StringVar(&cfg.DB.Sql.Name, "sql-name", sqlName, "SQL Database Name")

	flag.StringVar(&cfgFile, "c", cfgFile, "Config file path")

	flag.Parse()

	// reading settings from file, if applicable
	err := readConfigs(&cfg, cfgFile)
	if err != nil {
		return nil, err
	}

	err = verifyConfigs(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readConfigs(cfg *config.Config, cfgFile string) error {
	// no config file was provided
	if cfgFile == "" {
		return nil
	}

	file, err := os.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func verifyConfigs(cfg *config.Config) error {
	if cfg.Environment == "" {
		return fmt.Errorf("environment must be specified")
	}
	if !(cfg.Environment == config.Dev || cfg.Environment == config.Prod) {
		return fmt.Errorf("invalid environment, options are: %s - %s", config.Dev, config.Prod)
	}

	if cfg.Port == "" {
		return fmt.Errorf("port must be specified")
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 0 || port > 65535 {
		return fmt.Errorf("port must be a number between 0 and 65535")
	}

	if cfg.JWT.Secret == "" {
		return fmt.Errorf("JWT secret must be specified")
	}
	if cfg.JWT.Expire == "" {
		cfg.JWT.Expire = "0"
	}
	_, err = time.ParseDuration(cfg.JWT.Expire)
	if err != nil {
		return fmt.Errorf("JWT expire must be a valid duration, %v", err)
	}

	if cfg.DB.Sql.Host == "" {
		return fmt.Errorf("sql host must be specified")
	}
	if cfg.DB.Sql.Port == "" {
		return fmt.Errorf("sql port must be specified")
	}
	if cfg.DB.Sql.Name == "" {
		return fmt.Errorf("sql name must be specified")
	}
	if cfg.DB.Sql.Username == "" {
		return fmt.Errorf("sql username must be specified")
	}

	if cfg.Logger.Level == "" {
		cfg.Logger.Level = config.Info
	}
	if !(cfg.Logger.Level == config.Debug ||
		cfg.Logger.Level == config.Info ||
		cfg.Logger.Level == config.Warn ||
		cfg.Logger.Level == config.Error) {
		return fmt.Errorf(
			"invalid log level, options are: %s - %s - %s - %s",
			config.Debug,
			config.Info,
			config.Warn,
			config.Error,
		)
	}

	return nil
}

func getLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil

	case "info":
		return slog.LevelInfo, nil

	case "warn":
		return slog.LevelWarn, nil

	case "error":
		return slog.LevelError, nil

	default:
		return 0, fmt.Errorf("invalid log level, %s", level)
	}
}
