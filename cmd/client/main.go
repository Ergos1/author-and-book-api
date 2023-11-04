package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/route256/workshop-8/pkg/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := pb.NewMessagesClient(conn)

	summary, err := client.GetMessagesSummary(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	log.Printf("count %d", summary.GetCount())
	return nil
}
