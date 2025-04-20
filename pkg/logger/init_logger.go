package logger

import (
	"log/slog"
	"os"
)

func InitLogger(env string) *slog.Logger {
	var handler slog.Handler

	switch env {
	case "local":
		handler = slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		)
	case "prod":
		handler = slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{
				Level:     slog.LevelInfo,
				AddSource: true,
			},
		)
	}

	logger := slog.New(handler)

	return logger
}
