package proxy

import (
	"net/http"
)

type Proxy interface {
	ServeHttp(w http.ResponseWriter, r *http.Request, target string) error
}

func New() Proxy {
	return newReverse()
}
