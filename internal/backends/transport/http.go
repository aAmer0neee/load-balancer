package transport

import (
	"net/http"
)

type Http struct {
	Srv *http.Server
}

// просто заглушка, которая возвращает ответ
func NewHttpHandler(addr string)Http{
	return Http{
		Srv: &http.Server{
				Addr: addr,
				Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("Hello World!\nFrom " + addr))
				}),
		},
	}
}