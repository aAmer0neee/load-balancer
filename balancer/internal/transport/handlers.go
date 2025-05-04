package transport

import (
	"encoding/json"
	"net/http"
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

func configureHandlers(h *Http) {

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

	h.Srv.Handler = mux
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
