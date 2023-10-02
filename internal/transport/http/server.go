package http

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	Address string
	ctx     context.Context
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx: ctx,
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) basicHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("HI")
	})
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.Address,
		Handler: s.basicHandler(),
	}

	go s.Stop(srv)

	log.Println("[HTTP] server runing on: ", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) Stop(srv *http.Server) {
	<-s.ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] shutdown with error %v", err)
	}

	log.Print("[HTTP] shutdown")
}
