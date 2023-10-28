package http

import (
	"context"
	"net/http"

	"github.com/Shopify/sarama"
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

type KafkaProducer interface {
	SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type KafkaConsumer interface {
	Subscribe(ctx context.Context, topic string, handler func(message *sarama.ConsumerMessage)) error
}

func WithKafkaProducer(producer KafkaProducer) ServerOption {
	return func(srv *Server) {
		srv.Producer = producer
	}
}

func WithKafkaConsumer(consumer KafkaConsumer) ServerOption {
	return func(srv *Server) {
		srv.Consumer = consumer
	}
}
