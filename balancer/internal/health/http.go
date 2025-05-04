package health

import (
	"net/http"
	"time"

	"github.com/aAmer0neee/load-balancer/balancer/domain"
)

type HttpCheck struct {
	Timeout time.Duration
}

func (h *HttpCheck) Check(target string) bool {
	client := http.Client{Timeout: h.Timeout}
	resp, err := client.Get("http://" + target)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func newHttp(cfg domain.Cfg) *HttpCheck {
	return &HttpCheck{
		Timeout: time.Duration(cfg.Health.Timeout) * time.Millisecond,
	}
}
