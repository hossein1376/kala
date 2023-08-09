package Logger

import (
	"io"
	"log/slog"
)

func NewJsonLogger(dst io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(dst, nil)
	logger := slog.New(jsonHandler)
	return logger
}

func NewTextLogger(dst io.Writer) *slog.Logger {
	textLogger := slog.NewTextHandler(dst, nil)
	logger := slog.New(textLogger)
	return logger
}
