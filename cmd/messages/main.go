package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/route256/workshop-8/pkg/messages"
	"google.golang.org/grpc"

	impl "github.com/route256/workshop-8/internal/app/messages"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "add", ":50051", "Add for messages server")

	if err := run(ctx, addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
	server := grpc.NewServer()

	pb.RegisterMessagesServer(server, impl.New())

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("service messages listening on %q", addr)
	return server.Serve(lis)
}
