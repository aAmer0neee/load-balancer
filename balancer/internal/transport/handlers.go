package transport

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/aAmer0neee/load-balancer/balancer/internal/limiter"
)

type ResponseError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ResponseOK struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func configureHandlers(h *Http)http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusBadRequest, ResponseError{
				Code:  http.StatusBadRequest,
				Error: "Bad method",
			})
		} else {
			h.Hello(w, r)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusBadRequest, ResponseError{
				Code:  http.StatusBadRequest,
				Error: "Bad method",
			})
		} else {
			h.Hello(w, r)
		}
	})
	

	return mux
}

func logRequest(next http.Handler , log *slog.Logger)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/favicon.ico" {
			return
		}
		log.Info("New request","TARGET ADDR",r.RequestURI, "METHOD", r.Method, "ADDR", r.RemoteAddr)
		next.ServeHTTP(w,r)
	})
}

func limitMiddleware(next http.Handler, l limiter.Limiter)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		err := l.TakeToken(ip)
		if errors.Is(err, limiter.ErrLimitExceeded) {
			writeJSON(w, http.StatusTooManyRequests,ResponseError{
				Code: http.StatusTooManyRequests,
				Error: "To many Requests",
			})
			return
		} 
		next.ServeHTTP(w,r)

	})
}

func (h *Http) Hello(w http.ResponseWriter, r *http.Request) {

	h.service.HandleRequest(w, r)
	// writeJSON(w, http.StatusOK, ResponseOK{
	// 	Code:    http.StatusOK,
	// 	Message: "OK",
	// 	Data:    "Hello World!",
	// })
}

func writeJSON(w http.ResponseWriter, code int, payload any) {
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
