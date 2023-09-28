package Logger

import (
	"fmt"
	"io"
	"log/slog"
)

func NewJsonLogger(dst io.Writer, level string) *slog.Logger {
	slogLevel := getLevel(level)

	jsonHandler := slog.NewJSONHandler(dst, &slog.HandlerOptions{Level: slogLevel})
	logger := slog.New(jsonHandler)
	return logger
}

func NewTextLogger(dst io.Writer, level string) *slog.Logger {
	slogLevel := getLevel(level)

	textLogger := slog.NewTextHandler(dst, &slog.HandlerOptions{Level: slogLevel})
	logger := slog.New(textLogger)
	return logger
}

func getLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug

	case "info":
		return slog.LevelInfo

	case "warn":
		return slog.LevelWarn

	case "error":
		return slog.LevelError

	default:
		panic(fmt.Sprintf("invalid log level, %s", level))
	}
}
