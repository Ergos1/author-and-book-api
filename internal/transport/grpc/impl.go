package grpc

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	author_pb "gitlab.ozon.dev/ergossteam/homework-3/pkg/api/grpc/v1/author"
	book_pb "gitlab.ozon.dev/ergossteam/homework-3/pkg/api/grpc/v1/book"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GetAuthorById(ctx context.Context, id int64) (*core.AuthorWithBooks, error)
	CreateAuthor(ctx context.Context, request core.CreateAuthorRequest) (int64, error)
	UpdateAuthor(ctx context.Context, request core.UpdateAuthorRequest) error
	DeleteAuthorById(ctx context.Context, id int64) error
}

type Implementation struct {
	author_pb.UnimplementedAuthorServiceServer

	service Service
}

func New(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}

func mapToPBBook(book *book.Book) *book_pb.Book {
	return &book_pb.Book{
		Id:       book.ID,
		Rating:   book.Rating,
		AuthorId: book.AuthorID,
		Name:     book.Name,
	}
}

func mapToPBAuthor(author *core.AuthorWithBooks) *author_pb.Author {
	pb_books := make([]*book_pb.Book, 0, len(author.Books))
	for _, book := range author.Books {
		pb_books = append(pb_books, mapToPBBook(book))
	}

	return &author_pb.Author{
		Id:    author.ID,
		Name:  author.Name,
		Books: pb_books,
	}
}

func (i *Implementation) GetByID(ctx context.Context, req *author_pb.GetByIDRequest) (*author_pb.GetByIDResponse, error) {
	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("component", "grpc_impl")))

	logger.Infof(ctx, "author get by id called")

	author, err := i.service.GetAuthorById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	result := &author_pb.GetByIDResponse{
		Author: mapToPBAuthor(author),
	}

	return result, nil
}

func (i *Implementation) Create(ctx context.Context, req *author_pb.CreateRequest) (*author_pb.CreateReponse, error) {
	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("component", "grpc_impl")))

	logger.Infof(ctx, "author create with id called")

	createdAuthorID, err := i.service.CreateAuthor(ctx, core.CreateAuthorRequest{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	result := &author_pb.CreateReponse{
		Id: createdAuthorID,
	}

	return result, nil
}

func (i *Implementation) Update(ctx context.Context, req *author_pb.UpdateRequest) (*author_pb.UpdateResponse, error) {
	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("component", "grpc_impl")))

	logger.Infof(ctx, "author update by id called")

	err := i.service.UpdateAuthor(ctx, core.UpdateAuthorRequest{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &author_pb.UpdateResponse{}, nil
}

func (i *Implementation) Delete(ctx context.Context, req *author_pb.DeleteRequest) (*author_pb.DeleteResponse, error) {
	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("component", "grpc_impl")))

	logger.Infof(ctx, "author delete by id called")

	err := i.service.DeleteAuthorById(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &author_pb.DeleteResponse{}, nil
}
