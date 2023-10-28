package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ParseID(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
