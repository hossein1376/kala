package handlers

import (
	"net/http"
	"strconv"

	"kala/cmd"
	"kala/pkg/Errors"
	"kala/pkg/Json"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	app   *cmd.Application
	error *Errors.Errors
	Json.Json
}

func NewHandlers(app *cmd.Application) *handler {
	return &handler{
		app:   app,
		error: Errors.NewErrors(app.Logger),
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
