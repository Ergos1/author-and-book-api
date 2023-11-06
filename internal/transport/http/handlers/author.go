package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	author_service "gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/dtos"
)

type AuthorHandler struct {
	service Service
}

func NewAuthorHandler(s Service) *AuthorHandler {
	return &AuthorHandler{
		service: s,
	}
}

func (h *AuthorHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", h.GetAuthorByID)
	r.Post("/", h.CreateAuthor)
	r.Put("/{id}", h.UpdateAuthor)
	r.Delete("/{id}", h.DeleteAuthor)

	return r
}

func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseID(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
		return
	}

	author, err := h.service.GetAuthorById(r.Context(), id)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, Response{Err: ErrAuthorNotFound.Error()})
		return
	}

	authorDto := dtos.ReadAuthorDTO{}
	authorDto.MapFromAuthorWithBooks(author)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{Data: authorDto})
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var createAuthorDTO dtos.CreateAuthorDTO
	if err := render.DecodeJSON(r.Body, &createAuthorDTO); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
		return
	}

	err := validator.New().Struct(createAuthorDTO)
	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, Response{Err: err.Error()})
		return
	}

	author := &core.CreateAuthorRequest{
		ID:   createAuthorDTO.Id,
		Name: createAuthorDTO.Name,
	}

	_, err = h.service.CreateAuthor(r.Context(), *author)
	if err != nil {
		if errors.Is(err, author_service.ErrAuthorDuplicate) {
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, Response{Err: err.Error()})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, Response{Err: err.Error()})
		}
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, Response{Data: author})
}

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := ParseID(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
		return
	}

	var updateAuthorDTO dtos.UpdateAuthorDTO
	if err := render.DecodeJSON(r.Body, &updateAuthorDTO); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
		return
	}

	err = validator.New().Struct(updateAuthorDTO)
	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, Response{Err: err.Error()})
		return
	}

	author := &core.UpdateAuthorRequest{
		ID:   id,
		Name: updateAuthorDTO.Name,
	}

	err = h.service.UpdateAuthor(r.Context(), *author)
	if err != nil {
		if errors.Is(err, author_service.ErrAuthorNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, Response{Err: err.Error()})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, Response{Err: err.Error()})
		}
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{Data: author})
}

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := ParseID(r)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
		return
	}

	err = h.service.DeleteAuthorById(r.Context(), id)
	if err != nil {
		if errors.Is(err, author_service.ErrAuthorNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, Response{Err: err.Error()})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, Response{Err: err.Error()})
		}
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{})
}
