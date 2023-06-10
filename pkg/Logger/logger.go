package Logger

import (
	"os"

	"golang.org/x/exp/slog"
)

func NewJsonLogger() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	return logger
}

func NewTextLogger() *slog.Logger {
	textLogger := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(textLogger)
	return logger
}
