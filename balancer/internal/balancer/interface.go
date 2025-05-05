package balancer

import (
	"errors"
	"log/slog"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
)

var (
	ErrAllDead = errors.New("All servers are dead")
)

type Balancer interface {
	Next() (string, error)
	UpdateHealth(target string, health bool)
	All() []string
}

func New(l *slog.Logger, cfg domain.Cfg) Balancer {
	return newRoundRobbin(cfg.Services.Pool)
}
