package balancer

import (
	"net/url"
	"sync"
)

type RoundRobbin struct {
	pool []*backend
	mu   sync.Mutex
	index int
}

type backend struct {
	url *url.URL
	alive bool
}

func newRoundRobbin(pool []string) *RoundRobbin {
	p := []*backend{}
	for _, rawURL := range pool {
		b := backend{}
		u, err := url.Parse("http://" + rawURL)

		if err != nil {
			b.alive = false
		} else {
			b.alive = true
		}
		b.url = u
		p = append(p, &b)
	}
	return &RoundRobbin{
		pool: p,
		mu: sync.Mutex{},
		index: 0,
	}
}

func (r *RoundRobbin) Next()*url.URL{
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < len(r.pool); i++{
		idx := (i + r.index) % len(r.pool)
		if r.pool[idx].alive {
			r.index = (idx + 1) %len(r.pool)
			return r.pool[idx].url
		}
	}
	return nil
}
