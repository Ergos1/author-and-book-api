package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	book_service "gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.Check)

	return r
}

func (h *BaseHandler) Check(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{Data: "ok"})
}

type Service interface {
	GetAuthorById(ctx context.Context, id int64) (*core.AuthorWithBooks, error)
	CreateAuthor(ctx context.Context, request core.CreateAuthorRequest) (int64, error)
	GetBooksByAuthorID(ctx context.Context, authorID int64) ([]*book_service.Book, error)
	GetBookById(ctx context.Context, id int64) (*book_service.Book, error)
	CreateBook(ctx context.Context, request core.CreateBookRequest) (int64, error)
}
