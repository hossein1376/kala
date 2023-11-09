package logger

import (
	"io"
	"log/slog"
)

func NewJsonLogger(dst io.Writer, level slog.Level) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(dst, &slog.HandlerOptions{Level: level})
	logger := slog.New(jsonHandler)
	return logger
}

func NewTextLogger(dst io.Writer, level slog.Level) *slog.Logger {
	textLogger := slog.NewTextHandler(dst, &slog.HandlerOptions{Level: level})
	logger := slog.New(textLogger)
	return logger
}
