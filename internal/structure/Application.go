package structure

import (
	"database/sql"

	"golang.org/x/exp/slog"
)

type Application struct {
	Config Config
	Logger *slog.Logger
	Sql    *sql.DB
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
