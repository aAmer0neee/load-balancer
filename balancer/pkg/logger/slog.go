package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					t := a.Value.Time()
					return slog.String("time", t.Format("2006-01-02 15:04:05"))
				}
				return a
			},
		}))
}
