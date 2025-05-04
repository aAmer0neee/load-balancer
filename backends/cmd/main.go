package main

import (
	"github.com/aAmer0neee/load-balancer/backends/internal/transport"
	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/pkg/config"
	"github.com/aAmer0neee/load-balancer/balancer/pkg/logger"
)

func main() {

	cfg := domain.Cfg{}
	config.MustLoad(&cfg)

	log := logger.New()

	for _, addr := range cfg.Services.Pool {
		go func(addr string) {
			r := transport.NewHttpHandler(addr)
			log.Info("BACKEND SERVICE STARTING:", "ADDR", addr)
			if err := r.Srv.ListenAndServe(); err != nil {
				log.Warn("ERROR START BACKEND SERVER", "ADDR", addr)
			}
		}(addr)
	}
	select {}
}
