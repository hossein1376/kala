package Logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var Logger = InitLogger()

func InitLogger() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	return logger
}
