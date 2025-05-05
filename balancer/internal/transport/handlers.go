package transport

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strconv"

	"github.com/aAmer0neee/load-balancer/balancer/internal/limiter"
	"github.com/aAmer0neee/load-balancer/balancer/internal/service"
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

func (h *Http)configureHandlers()http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			h.writeJSON(w, http.StatusBadRequest, ResponseError{
				Code:  http.StatusBadRequest,
				Error: "Bad method",
			})
		} else {
			h.Hello(w, r)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			h.writeJSON(w, http.StatusBadRequest, ResponseError{
				Code:  http.StatusBadRequest,
				Error: "Bad method",
			})
		} else {
			h.Hello(w, r)
		}
	})
	

	return mux
}

func (h *Http)limitMiddleware(next http.Handler, l limiter.Limiter)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		err := l.TakeToken(ip)
		if errors.Is(err, limiter.ErrLimitExceeded) {
			h.writeJSON(w, http.StatusTooManyRequests,ResponseError{
				Code: http.StatusTooManyRequests,
				Error: "To many Requests",
			})
			return
		} 
		next.ServeHTTP(w,r)

	})
}

func (h *Http) Hello(w http.ResponseWriter, r *http.Request) {

	if err := h.service.HandleRequest(w, r); errors.Is(err, service.ErrInternalService) {
		h.writeJSON(w, http.StatusInternalServerError, ResponseError{
			Code: http.StatusInternalServerError,
			Error: "Internal service error",
		})
	}
	
	// writeJSON(w, http.StatusOK, ResponseOK{
	// 	Code:    http.StatusOK,
	// 	Message: "OK",
	// 	Data:    "Hello World!",
	// })
}

// логгер исходящих запросо + отправляет JSON ответ
func (h *Http)writeJSON(w http.ResponseWriter, code int, payload any) {
	h.log.Info("response",strconv.Itoa(code) ,http.StatusText(code))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// логгер входящих запросов
func (h *Http)logRequest(next http.Handler )http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/favicon.ico" {
			return
		}
		h.log.Info("New request","TARGET ADDR",r.RequestURI, "METHOD", r.Method, "ADDR", r.RemoteAddr)
		next.ServeHTTP(w,r)
	})
}
