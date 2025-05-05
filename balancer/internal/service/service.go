package service

import (
	"errors"

	"log/slog"
	"net/http"

	"github.com/aAmer0neee/load-balancer/balancer/internal/balancer"
	"github.com/aAmer0neee/load-balancer/balancer/internal/proxy"
)

var (
	ErrInternalService = errors.New("Internal Service Error")
)

type GatewayService interface {
	HandleRequest(w http.ResponseWriter, r *http.Request) error
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

// основной сервис, запрашивает живой сервис, и проксирует, в случае ошибки пробует несколько раз
func (g *Gateway) HandleRequest(w http.ResponseWriter, r *http.Request) error {

	maxAttemp := 5
	for i := 0; i <= maxAttemp; i++ {
		backendURL, err := g.balancer.Next()
		if err != nil {
			continue
		}
		err = g.proxy.ServeHttp(w, r, backendURL)
		if err == nil {
			g.log.Info("Forwarding request", "proxy", backendURL)
			return nil
		} else {
			g.log.Warn("Error send request", "proxy", backendURL, "message", err.Error())
			continue
		}

	}
	g.log.Warn("no alive servers")
	return ErrInternalService

}
