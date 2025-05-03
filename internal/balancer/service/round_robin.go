package service

import "log/slog"

type RoundRobbin struct {
	log *slog.Logger
}

func newRoundRobbin(l *slog.Logger) *RoundRobbin {
	return &RoundRobbin{log: l}
}

func (r *RoundRobbin) Forward() {

}
