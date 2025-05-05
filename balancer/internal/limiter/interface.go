package limiter

import (
	"errors"


	"github.com/aAmer0neee/load-balancer/balancer/domain"

)

var (
	ErrNotFound      = errors.New("bucket for user not found")
	ErrLimitExceeded = errors.New("limit for user exceeded")
)

type Limiter interface {
	// следует заменить на UUID 
	// и вынести логику обработки уникальных клиентов наружу, чтобы мидлвейр ничего не знал об ip
	TakeToken(id string) error
	StartRefillTokens()
}

func New(cfg domain.Cfg) Limiter {
	return newBucket(cfg.Limiter.Ticker, cfg.Limiter.Capacity)
}
