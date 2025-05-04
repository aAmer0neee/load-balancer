package service

import (
	"log/slog"
	"net/http"

	"github.com/aAmer0neee/load-balancer/balancer/internal/balancer"
	"github.com/aAmer0neee/load-balancer/balancer/internal/proxy"
)

type GatewayService interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

func New(l *slog.Logger, p proxy.Proxy, b balancer.Balancer)GatewayService{
	return &Gateway{
		log: l,
		balancer: b,
		proxy: p,
	}
}

type Gateway struct {
	log *slog.Logger
	balancer balancer.Balancer
	proxy proxy.Proxy

}

func (g *Gateway)HandleRequest(w http.ResponseWriter, r *http.Request){
	backendURL := g.balancer.Next()

	g.proxy.ServeHttp(w,r,backendURL)
	g.log.Debug("Proxy request to", "server url", backendURL)
}