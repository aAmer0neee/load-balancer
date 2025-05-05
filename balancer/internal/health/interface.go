package health

import (
	"log/slog"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/internal/balancer"
)

type Health interface {
	Check(target string) bool
	HealthMonitoring(period int, b balancer.Balancer, log *slog.Logger)
}

func New(cfg domain.Cfg) Health {
	return newHttp(cfg)
}
