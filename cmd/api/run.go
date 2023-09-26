package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-chi/jwtauth/v5"

	"github.com/hossein1376/kala/config"
	"github.com/hossein1376/kala/internal/data"
	"github.com/hossein1376/kala/internal/handlers"
	"github.com/hossein1376/kala/pkg/Logger"
)

func runServer() {
	logger := Logger.NewJsonLogger(os.Stdout)

	cfg, err := newConfig()
	if err != nil {
		logger.Error("failed to load configurations", slog.String("error", err.Error()))
		return
	}
	logger.Info("configurations loaded successfully")

	client, err := openSqlDB(cfg)
	if err != nil {
		logger.Error("failed to open sql connection", slog.String("error", err.Error()))
		return
	}
	defer client.Close()
	logger.Info("sql database connection established")

	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Error("failed creating schema resources: %v", slog.String("error", err.Error()))
		return
	}

	cfg.JWTToken = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	logger.Info("JWT token generated")

	app := &config.Application{
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
