package http

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
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
				var body any
				if err := render.DecodeJSON(r.Body, &body); err != nil {
					body = nil
				}

				if err := requestLogger.Log(r.Method, r.URL.Path, body); err != nil {
					log.Printf("kafka error: %v", err)
				}
			}
			return http.HandlerFunc(fn)
		})
	}
}
