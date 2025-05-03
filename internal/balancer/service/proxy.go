package service

import (
	"net/http/httputil"
	"net/url"
)

func Proxy(target *url.URL){
	httputil.NewSingleHostReverseProxy(target)
}