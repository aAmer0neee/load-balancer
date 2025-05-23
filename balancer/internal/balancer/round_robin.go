package balancer

import (
	"sync"
)

type RoundRobbin struct {
	pool  []*backend
	mu    sync.Mutex
	index int
}

type backend struct {
	url   string
	alive bool
}

func newRoundRobbin(pool []string) *RoundRobbin {
	p := []*backend{}
	for _, rawURL := range pool {
		b := backend{}
		b.alive = true
		b.url = rawURL
		p = append(p, &b)
	}
	return &RoundRobbin{
		pool:  p,
		mu:    sync.Mutex{},
		index: 0,
	}
}

// возвращает живой сервер, если такого нет, то возвращает ошибку 
func (r *RoundRobbin) Next() (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < len(r.pool); i++ {
		idx := (i + r.index) % len(r.pool)
		if r.pool[idx].alive {
			r.index = (idx + 1) % len(r.pool)
			return r.pool[idx].url, nil
		}
	}
	return "", ErrAllDead
}

// обновляет здоровье сервера
func (r *RoundRobbin) UpdateHealth(target string, health bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, server := range r.pool {
		if server.url == target {
			server.alive = health
			break
		}
	}
}

// возвращает список серверов, для периодичной проврки состояния
func (r *RoundRobbin) All() []string {
	all := []string{}
	for _, server := range r.pool {
		all = append(all, server.url)
	}
	return all
}
