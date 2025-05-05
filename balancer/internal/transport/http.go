package transport

import (
	"log/slog"
	"net/http"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/internal/limiter"
	"github.com/aAmer0neee/load-balancer/balancer/internal/service"
)

type Http struct {
	Srv     http.Server
	service service.GatewayService
}

func NewHttpHandler(s service.GatewayService, cfg domain.Cfg, l limiter.Limiter, log *slog.Logger) *Http {
	h := &Http{
		service: s,
		Srv: http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		},
	}

	router := configureHandlers(h)

	handler := limitMiddleware(router, l)


	h.Srv.Handler = logRequest(handler, log) 
	return h
}
