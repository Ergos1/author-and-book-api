package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/route256/workshop-8/pkg/gateway"
	messages_pb "github.com/route256/workshop-8/pkg/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	impl "github.com/route256/workshop-8/internal/app/gateway"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "add", ":50052", "Add for messages server")

	if err := run(ctx, addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
	conn, err := grpc.Dial(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	gw := impl.New(messages_pb.NewMessagesClient(conn))

	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, gw)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("service gateway listening on %q", addr)
	return server.Serve(lis)
}
