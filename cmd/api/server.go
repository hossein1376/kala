package api

import (
	"database/sql"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

type application struct {
	config config
	logger *slog.Logger
	sql    *sql.DB
}

func RunServer() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	cfg := newConfig()
	db, err := openSqlDB(cfg)
	if err != nil {
		logger.Error("failed to open sql connection", slog.String("error", err.Error()))
		return
	}
	defer db.Close()
	logger.Info("database connection established")

	app := application{
		config: cfg,
		logger: logger,
		sql:    db,
	}

	fmt.Println(app) // just to get it to compile for now..
}
