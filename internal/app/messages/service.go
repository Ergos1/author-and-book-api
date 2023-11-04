package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

func (i *Implementation) PullMessages(_ *emptypb.Empty, stream pb.Messages_PullMessagesServer) error {
	for _, message := range data {
		stream.Send(message)
	}
	return nil
}

func (i *Implementation) PushMessages(stream pb.Messages_PushMessagesServer) error {
	var counter int
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.MessagesSummary{
					Count: uint64(counter),
				})
			}
			return err
		}

		log.Printf("message %q from %s", msg.Text, msg.Author)
		counter++
	}
}

func (i *Implementation) ExchangeMessages(stream pb.Messages_ExchangeMessagesServer) error {
	for {
		message, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		stream.Send(&pb.Message{
			Ts:     time.Now().Format(time.RFC3339),
			Text:   fmt.Sprintf("RE: %s", message.Text),
			Author: "server",
		})

		log.Printf("message %q from %s", message.Text, message.Author)
	}
}

var data = []*pb.Message{
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 1",
		Author: "server",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 2",
		Author: "server",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 3",
		Author: "server",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 4",
		Author: "server",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 5",
		Author: "server",
	},
}
