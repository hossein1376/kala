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
	Sql         SqlDB
	Environment string
	Port        string
	JWTSecret   string
	JWTToken    *jwtauth.JWTAuth
}

type SqlDB struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}
