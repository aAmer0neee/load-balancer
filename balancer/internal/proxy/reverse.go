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

// проксирует реквест на указаный сервер
func (p *ReverseProxy) ServeHttp(w http.ResponseWriter, r *http.Request, target string) error {
	url, err := url.Parse("http://" + target)
	if err != nil {
		return err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
	return nil
}
