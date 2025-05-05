package health

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
	"github.com/aAmer0neee/load-balancer/balancer/internal/balancer"
)

type HttpCheck struct {
	Timeout time.Duration
}

// проверяет один сервер
func (h *HttpCheck) Check(target string) bool {
	client := http.Client{Timeout: h.Timeout}
	resp, err := client.Get("http://" + target)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// периодичный чек всех серверов
func (h *HttpCheck) HealthMonitoring(period int, b balancer.Balancer, log *slog.Logger) {
	go func() {
		ticker := time.NewTicker(time.Duration(period) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {

			for _, server := range b.All() {
				go func(url string) {
					status := h.Check(url)
					if !status {
						log.Warn("Server dead", "server", url, "alive", status)
					}
					b.UpdateHealth(url, status)
				}(server)

			}
		}
	}()
}

func newHttp(cfg domain.Cfg) *HttpCheck {
	return &HttpCheck{
		Timeout: time.Duration(cfg.Health.Timeout) * time.Millisecond,
	}
}
