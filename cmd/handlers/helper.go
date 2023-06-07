package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"kala/cmd"
)

type Handler struct {
	app cmd.Application
}

func NewHandlers(app cmd.Application) Handler {
	return Handler{
		app: app,
	}
}

// paramInt extracts ID as an integer value from the URL parameter.
func paramInt(r *http.Request, name string) int {
	id, err := strconv.Atoi(chi.URLParam(r, name))
	if err != nil {
		return 0
	}
	return id
}

func insertLogs(route, method string, id int) error {

	return nil
}
