package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// paramInt extracts ID as an integer value from the URL parameter.
func paramInt(r *http.Request) int {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return 0
	}
	return id
}
