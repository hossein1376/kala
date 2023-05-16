package api

import (
	"golang.org/x/exp/slog"
	"kala/cmd/handlers"
	"kala/internal/Logger"
	"kala/internal/data"
	"kala/internal/structure"
)

func RunServer() {
	logger := Logger.Logger

	cfg := newConfig()
	client, err := openSqlDB(cfg)
	if err != nil {
		logger.Error("failed to open sql connection", slog.String("error", err.Error()))
		return
	}
	defer client.Close()
	logger.Info("database connection established")

	app := structure.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(client),
	}

	router := handlers.NewHandler(app)
	err = serve(app, router)
	if err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		return
	}
}
