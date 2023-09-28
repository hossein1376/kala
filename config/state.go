package config

import (
	"log/slog"

	"github.com/go-chi/jwtauth/v5"

	"github.com/hossein1376/kala/internal/data"
)

const (
	Dev  = "dev"
	Prod = "prod"

	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

type Application struct {
	Config *Config
	Logger *slog.Logger
	Models *data.Models
}

type Config struct {
	Environment string `json:"environment"`
	Port        string `json:"port"`
	DB          DB     `json:"db"`
	JWT         JWT    `json:"jwt"`
	Logger      Logger `json:"logger"`
}

type JWT struct {
	Secret string           `json:"secret"`
	Expire string           `json:"expire"`
	Token  *jwtauth.JWTAuth `json:"-"`
}

type Logger struct {
	Level string `json:"level"`
}

type DB struct {
	Sql   Sql   `json:"sql"`
	NoSql NoSql `json:"no_sql"`
	Redis Redis `json:"redis"`
}

type Sql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type NoSql struct {
}

type Redis struct {
}
