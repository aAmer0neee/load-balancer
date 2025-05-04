package transport

import (
	"net/http"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/internal/service"
)

type Http struct {
	Srv     http.Server
	service service.GatewayService
}

func NewHttpHandler(s service.GatewayService, cfg domain.Cfg) *Http {
	h := &Http{
		service: s,
		Srv: http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		},
	}

	configureHandlers(h)

	return h
}
