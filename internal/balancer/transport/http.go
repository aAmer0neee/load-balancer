package transport

import (
	"net/http"

	"github.com/aAmer0neee/load-balancer/domain"
	"github.com/aAmer0neee/load-balancer/internal/balancer/service"
)

type Http struct {
	Srv     http.Server
	service service.Balancer
}

func NewHttpHandler(s service.Balancer, cfg domain.Cfg) *Http {
	h := &Http{service: s,
		Srv: http.Server{
			Addr: cfg.Server.Host + ":" + cfg.Server.Port,
		},
	}

	configureHandlers(h)

	return h
}
