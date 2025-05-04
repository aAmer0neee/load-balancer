package proxy

import (
	"net/http"
	"net/url"
)

type Proxy interface {
	ServeHttp(w http.ResponseWriter, r *http.Request, target *url.URL) error
}

func New() Proxy {
	return newReverse()
}
