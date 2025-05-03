package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aAmer0neee/load-balancer/domain"
	"github.com/aAmer0neee/load-balancer/internal/balancer/service"
	"github.com/aAmer0neee/load-balancer/internal/balancer/transport"
	"github.com/aAmer0neee/load-balancer/pkg/config"
	"github.com/aAmer0neee/load-balancer/pkg/logger"
)

func main() {
	cfg := domain.Cfg{}
	config.MustLoad(&cfg)

	log := logger.New()

	service := service.New(log)

	r := transport.NewHttpHandler(service, cfg)

	log.Info("SERVER STARTING", "HOST", cfg.Server.Host, "PORT", cfg.Server.Port)
	go func() {
		shutdown(r, log)
	}()
	if err := r.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Info("ERORR START SERVER", "message", err.Error())
		os.Exit(1)
	}

}

func shutdown(r *transport.Http, log *slog.Logger) {
	wait := make(chan os.Signal, 1)

	signal.Notify(wait, syscall.SIGTERM, os.Interrupt)

	<-wait

	log.Info("Graceful Shutdown", "WAIT", "7 Sec")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	r.Srv.Shutdown(ctx)
}
