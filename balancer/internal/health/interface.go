package health

import "github.com/aAmer0neee/load-balancer/balancer/domain"

type Health interface {
	Check(target string) bool
}

func New(cfg domain.Cfg) Health {
	return newHttp(cfg)
}
