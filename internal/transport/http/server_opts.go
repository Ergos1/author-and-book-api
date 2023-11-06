package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithMount(pattern string, handler http.Handler) ServerOption {
	return func(srv *Server) {
		srv.Handler.Mount(pattern, handler)
	}
}

type RequestLogger interface {
	Log(method string, url string, body any) error
}

func WithRequestLogger(requestLogger RequestLogger) ServerOption {
	return func(srv *Server) {
		srv.Handler.Use(func(next http.Handler) http.Handler {
			fn := func(w http.ResponseWriter, r *http.Request) {
				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Failed to read body", http.StatusInternalServerError)
					return
				}

				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				var body any
				if err := json.Unmarshal(bodyBytes, &body); err != nil {
					body = nil
				}

				if err := requestLogger.Log(r.Method, r.URL.Path, body); err != nil {
					log.Printf("kafka error: %v", err)
				}

				next.ServeHTTP(w, r)
			}
			return http.HandlerFunc(fn)
		})
	}
}
