package gateway

import (
	"context"

	pb "github.com/route256/workshop-8/pkg/gateway"
	pb_messages "github.com/route256/workshop-8/pkg/messages"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	pb.UnimplementedGatewayServer

	messages pb_messages.MessagesClient
}

func New(messages pb_messages.MessagesClient) *Implementation {
	return &Implementation{
		messages: messages,
	}
}

func (i *Implementation) GetMessagesSummary(ctx context.Context, empt *emptypb.Empty) (*pb.MessagesSummary, error) {
	summary, err := i.messages.GetMessagesSummary(ctx, empt)
	if err != nil {
		return nil, err
	}

	return &pb.MessagesSummary{
		Count: summary.GetCount(),
	}, nil
}
