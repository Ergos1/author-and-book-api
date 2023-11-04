package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/route256/workshop-8/pkg/messages"
)

type Implementation struct {
	pb.UnimplementedMessagesServer
}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) GetMessagesSummary(ctx context.Context, _ *emptypb.Empty) (*pb.MessagesSummary, error) {

	return &pb.MessagesSummary{Count: 42}, nil
}
