package main

import (
	"context"
	"log"
	"os"

	"github.com/go-chi/jwtauth/v5"

	"github.com/hossein1376/kala/config"
	"github.com/hossein1376/kala/internal/data"
	"github.com/hossein1376/kala/internal/handlers"
	"github.com/hossein1376/kala/pkg/Logger"
)

func runServer() {
	cfg, err := newConfig()
	if err != nil {
		log.Println("failed to load configurations,", err)
		return
	}

	logger := Logger.NewJsonLogger(os.Stdout, cfg.Logger.Level)
	logger.Info("configurations loaded successfully")

	client, err := openSqlDB(cfg)
	if err != nil {
		logger.Error("failed to open sql connection", "error", err)
		return
	}
	defer client.Close()
	logger.Info("sql database connection established")

	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Error("failed creating schema resources: %v", "error", err)
		return
	}

	cfg.JWT.Token = jwtauth.New("HS256", []byte(cfg.JWT.Secret), nil)
	logger.Info("JWT token initialized")

	app := &config.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(client),
	}

	handler := handlers.NewHandlers(app)
	router := handler.Router()
	err = serve(app, router)
	if err != nil {
		logger.Error("failed to start server", "error", err)
		return
	}
}
