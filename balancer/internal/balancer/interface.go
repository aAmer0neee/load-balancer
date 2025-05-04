package balancer

import (
	"log/slog"
	"net/url"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
)

type Balancer interface {
	Next()*url.URL
}

func New(l *slog.Logger, cfg domain.Cfg) Balancer {
	return newRoundRobbin(cfg.Services.Pool)
}
