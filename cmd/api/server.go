package api

import (
	"os"

	"golang.org/x/exp/slog"
	"kala/cmd/handlers"
	"kala/internal/structure"
)

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

	app := structure.Application{
		Config: cfg,
		Logger: logger,
		Sql:    db,
	}

	router := handlers.NewHandler(app)
	err = serve(app, router)
	if err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		return
	}
}