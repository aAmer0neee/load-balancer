package service

import "log/slog"

type Balancer interface {
	Forward()
}

func New(l *slog.Logger) Balancer {
	return newRoundRobbin(l)
}
