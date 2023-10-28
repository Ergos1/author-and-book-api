package http

import (
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
