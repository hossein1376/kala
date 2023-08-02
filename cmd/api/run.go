package api

import (
	"context"

	"github.com/hossein1376/kala/cmd"
	"github.com/hossein1376/kala/cmd/handlers"
	"github.com/hossein1376/kala/internal/data"
	"github.com/hossein1376/kala/pkg/Logger"

	"golang.org/x/exp/slog"
)

func RunServer() {
	logger := Logger.NewJsonLogger()

	cfg := newConfig()
	client, err := openSqlDB(cfg)
	if err != nil {
		logger.Error("failed to open sql connection", slog.String("error", err.Error()))
		return
	}
	defer client.Close()
	logger.Info("database connection established")

	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Error("failed creating schema resources: %v", slog.String("error", err.Error()))
		return
	}

	app := &cmd.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(client),
	}

	handler := handlers.NewHandlers(app)
	router := handler.Router()
	err = serve(app, router)
	if err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		return
	}
}
