package service

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/internal/balancer"
	"github.com/aAmer0neee/load-balancer/balancer/internal/health"
	"github.com/aAmer0neee/load-balancer/balancer/internal/proxy"
)

type GatewayService interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type Gateway struct {
	log      *slog.Logger
	balancer balancer.Balancer
	proxy    proxy.Proxy
}

func New(l *slog.Logger, p proxy.Proxy, b balancer.Balancer) GatewayService {
	return &Gateway{
		log:      l,
		balancer: b,
		proxy:    p,
	}
}

func (g *Gateway) HandleRequest(w http.ResponseWriter, r *http.Request) {

	backendURL := g.balancer.Next()

	g.proxy.ServeHttp(w, r, backendURL)
	g.log.Debug("Proxy request to", "server url", backendURL)
}

func StartMonitoring(cfg domain.Cfg, b balancer.Balancer, h health.Health) {

	go func() {
		ticker := time.NewTicker(time.Duration(cfg.Health.Ticker) * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {

			for _, server := range b.All() {
				go func(url string) {
					b.UpdateHealth(url, h.Check(url))
				}(server)
			}
		}
	}()
}
