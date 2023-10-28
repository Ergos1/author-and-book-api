package http

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	Address  string
	Handler  chi.Router
	ctx      context.Context
	Producer KafkaProducer
	Consumer KafkaConsumer
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx:     ctx,
		Handler: chi.NewRouter(),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.Address,
		Handler: s.Handler,
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
