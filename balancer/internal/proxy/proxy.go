package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct{}

func newReverse() Proxy {
	return &ReverseProxy{}
}

func (p *ReverseProxy) ServeHttp(w http.ResponseWriter, r *http.Request, target *url.URL) error {
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
	return nil
}
