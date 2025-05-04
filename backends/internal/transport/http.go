package transport

import (
	"encoding/json"
	"net/http"
)

type Http struct {
	Srv *http.Server
}

// просто заглушка, которая возвращает ответ
func NewHttpHandler(addr string) Http {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"health": "ok",
			"from":   addr,
		}

		json.NewEncoder(w).Encode(response)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Hello World!",
			"from":    addr,
		}
		json.NewEncoder(w).Encode(response)
	})
	return Http{
		Srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}
