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
	log *slog.Logger
}

func NewHttpHandler(s service.GatewayService, cfg domain.Cfg, l limiter.Limiter, log *slog.Logger) *Http {
	h := &Http{
		service: s,
		Srv: http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		},
		log: log,
	}

	router := h.configureHandlers()

	handler := h.limitMiddleware(router, l)


	h.Srv.Handler = h.logRequest(handler) 
	return h
}
