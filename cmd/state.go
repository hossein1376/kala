package cmd

import (
	"github.com/go-chi/jwtauth/v5"

	"github.com/hossein1376/kala/internal/data"

	"golang.org/x/exp/slog"
)

type Application struct {
	Config *Config
	Logger *slog.Logger
	Models *data.Models
}

type Config struct {
	Environment string
	Port        string
	JWTSecret   string
	JWTToken    *jwtauth.JWTAuth
	DB          DB
}

type DB struct {
	Sql   Sql
	NoSql NoSql
	Redis Redis
}

type Sql struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type NoSql struct {
}

type Redis struct {
}
