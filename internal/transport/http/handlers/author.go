package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	author_service "gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/dtos"
)

type RequestLogger interface {
	Log(method string, body any) error
}

type AuthorHandler struct {
	service       Service
	requestLogger RequestLogger
}

func NewAuthorHandler(s Service, requestLogger RequestLogger) *AuthorHandler {
	return &AuthorHandler{
		service:       s,
		requestLogger: requestLogger,
	}
}

func (h *AuthorHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", h.GetAuthorByID)
	r.Post("/", h.CreateAuthor)
	// r.Put("/{id}", ah.UpdateAuthor)
	// r.Delete("/{id}", ah.DeleteAuthor)

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

	if err := h.requestLogger.Log("GET", nil); err != nil {
		log.Printf("Request Logger Problem")
	}
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

	if err := h.requestLogger.Log("POST", createAuthorDTO); err != nil {
		log.Printf("Request Logger Problem")
	}
}

// func (ah *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
// 		return
// 	}

// 	var updateAuthorDTO dtos.UpdateAuthorDTO
// 	if err := render.DecodeJSON(r.Body, &updateAuthorDTO); err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
// 		return
// 	}

// 	err = validator.New().Struct(updateAuthorDTO)
// 	if err != nil {
// 		render.Status(r, http.StatusUnprocessableEntity)
// 		render.JSON(w, r, Response{Err: err.Error()})
// 		return
// 	}

// 	author := &models.Author{
// 		Id:   id,
// 		Name: updateAuthorDTO.Name,
// 	}

// 	err = ah.db.Authors().Update(r.Context(), id, author)
// 	if err != nil {
// 		if errors.Is(err, db.ErrObjectNotFound) {
// 			render.Status(r, http.StatusNotFound)
// 			render.JSON(w, r, Response{Err: ErrAuthorNotFound.Error()})
// 		} else {
// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		}
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// 	render.JSON(w, r, Response{Data: author})
// }

// func (ah *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
// 		return
// 	}

// 	err = ah.db.Authors().Delete(r.Context(), id)
// 	if err != nil {
// 		if errors.Is(err, db.ErrObjectNotFound) {
// 			render.Status(r, http.StatusNotFound)
// 			render.JSON(w, r, Response{Err: ErrAuthorNotFound.Error()})
// 		} else {
// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		}
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// 	render.JSON(w, r, Response{})
// }
