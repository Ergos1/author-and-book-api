package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	book_service "gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/dtos"
)

type BookHandler struct {
	service Service
}

func NewBookHandler(s Service) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", h.GetBookByID)
	r.Post("/", h.CreateBook)
	// r.Put("/{id}", ah.UpdateAuthor)
	// r.Delete("/{id}", ah.DeleteAuthor)

	return r
}

func (bh *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
		return
	}

	book, err := bh.service.GetBookById(r.Context(), id)
	if err != nil {
		if errors.Is(err, book_service.ErrBookNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, Response{Err: ErrBookNotFound.Error()})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, Response{Err: err.Error()})
		}
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{Data: book})
}

func (bh *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBookDTO dtos.CreateBookDTO
	if err := render.DecodeJSON(r.Body, &createBookDTO); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
		return
	}

	err := validator.New().Struct(createBookDTO)
	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, Response{Err: err.Error()})
		return
	}

	book := core.CreateBookRequest{
		ID:       createBookDTO.Id,
		Name:     createBookDTO.Name,
		Rating:   createBookDTO.Rating,
		AuthorID: createBookDTO.AuthorID,
	}

	_, err = bh.service.CreateBook(r.Context(), book)

	if err != nil {
		if errors.Is(err, book_service.ErrBookDuplicate) {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, Response{Err: err.Error()})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, Response{Err: err.Error()})
		}
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, Response{Data: book})
}

// func (bh *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
// 		return
// 	}

// 	var updateBookDTO dtos.UpdateBookDTO
// 	if err := render.DecodeJSON(r.Body, &updateBookDTO); err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
// 		return
// 	}

// 	err = validator.New().Struct(updateBookDTO)
// 	if err != nil {
// 		render.Status(r, http.StatusUnprocessableEntity)
// 		render.JSON(w, r, Response{Err: err.Error()})
// 		return
// 	}

// 	book := &models.Book{
// 		Id:       id,
// 		Name:     updateBookDTO.Name,
// 		Rating:   updateBookDTO.Rating,
// 		AuthorID: updateBookDTO.AuthorID,
// 	}

// 	err = bh.db.Books().Update(r.Context(), id, book)
// 	if err != nil {
// 		if errors.Is(err, db.ErrObjectNotFound) {
// 			render.Status(r, http.StatusNotFound)
// 			render.JSON(w, r, Response{Err: ErrBookNotFound.Error()})
// 		} else {
// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		}
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// 	render.JSON(w, r, Response{Data: book})
// }

// func (bh *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
// 		return
// 	}

// 	err = bh.db.Books().Delete(r.Context(), id)
// 	if err != nil {
// 		if errors.Is(err, db.ErrObjectNotFound) {
// 			render.Status(r, http.StatusNotFound)
// 			render.JSON(w, r, Response{Err: ErrBookNotFound.Error()})
// 		} else {
// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		}
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// 	render.JSON(w, r, Response{})
// }
