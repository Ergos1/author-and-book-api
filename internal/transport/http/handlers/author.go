package handlers

// import (
// 	"errors"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/render"
// 	"github.com/go-playground/validator/v10"
// 	"gitlab.ozon.dev/ergossteam/homework-3/internal/db"
// 	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
// 	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/dtos"
// )

// type AuthorHandler struct {
// 	db db.DB
// }

// func NewAuthorHandler(db db.DB) *AuthorHandler {
// 	return &AuthorHandler{
// 		db: db,
// 	}
// }

// func (ah *AuthorHandler) Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Get("/{id}", ah.GetAuthorByID)
// 	r.Post("/", ah.CreateAuthor)
// 	r.Put("/{id}", ah.UpdateAuthor)
// 	r.Delete("/{id}", ah.DeleteAuthor)

// 	return r
// }

// func (ah *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidId.Error()})
// 		return
// 	}

// 	author, err := ah.db.Authors().GetById(r.Context(), id)
// 	if err != nil {
// 		render.Status(r, http.StatusNotFound)
// 		render.JSON(w, r, Response{Err: ErrAuthorNotFound.Error()})
// 		return
// 	}

// 	books, err := ah.db.Books().GetByAuthorId(r.Context(), author.Id)
// 	if err != nil {
// 		render.Status(r, http.StatusInternalServerError)
// 		render.JSON(w, r, Response{Err: err.Error()})
// 		return
// 	}

// 	booksDto := make([]dtos.ReadBookDTO, 0, len(books))
// 	for _, book := range books {
// 		booksDto = append(booksDto, dtos.ReadBookDTO{
// 			Id:       book.Id,
// 			Name:     book.Name,
// 			Rating:   book.Rating,
// 			AuthorID: book.AuthorID,
// 		})
// 	}

// 	authorDto := dtos.ReadAuthorDTO{
// 		Id:    author.Id,
// 		Name:  author.Name,
// 		Books: booksDto,
// 	}

// 	render.Status(r, http.StatusOK)
// 	render.JSON(w, r, Response{Data: authorDto})
// }

// func (ah *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
// 	var createAuthorDTO dtos.CreateAuthorDTO
// 	if err := render.DecodeJSON(r.Body, &createAuthorDTO); err != nil {
// 		render.Status(r, http.StatusBadRequest)
// 		render.JSON(w, r, Response{Err: ErrInvalidRequestBody.Error()})
// 		return
// 	}

// 	err := validator.New().Struct(createAuthorDTO)
// 	if err != nil {
// 		render.Status(r, http.StatusUnprocessableEntity)
// 		render.JSON(w, r, Response{Err: err.Error()})
// 		return
// 	}

// 	author := &models.Author{
// 		Id:   createAuthorDTO.Id,
// 		Name: createAuthorDTO.Name,
// 	}

// 	_, err = ah.db.Authors().Create(r.Context(), author)
// 	if err != nil {
// 		if errors.Is(err, db.ErrDuplicate) {
// 			render.Status(r, http.StatusConflict)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		} else {
// 			render.Status(r, http.StatusInternalServerError)
// 			render.JSON(w, r, Response{Err: err.Error()})
// 		}
// 		return
// 	}

// 	render.Status(r, http.StatusCreated)
// 	render.JSON(w, r, Response{Data: author})
// }

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
