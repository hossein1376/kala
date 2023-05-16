package structure

import (
	"golang.org/x/exp/slog"
	"kala/internal/ent"
)

type Application struct {
	Config Config
	Logger *slog.Logger
	Client *ent.Client
}

type Config struct {
	Environment string
	Port        string
	Sql         SqlDB
}

type SqlDB struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}
