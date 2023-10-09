package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (ah *BaseHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ah.Check)

	return r
}

func (h *BaseHandler) Check(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, Response{Data: "ok"})
}
