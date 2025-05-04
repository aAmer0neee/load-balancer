package balancer

import (
	"log/slog"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
)

type Balancer interface {
	Next() string
	UpdateHealth(target string, health bool)
	All() []string
}

func New(l *slog.Logger, cfg domain.Cfg) Balancer {
	return newRoundRobbin(cfg.Services.Pool)
}
