package handlers

import (
	"net/http"
	"strconv"

	"github.com/hossein1376/kala/cmd"
	"github.com/hossein1376/kala/internal/errors"
	"github.com/hossein1376/kala/pkg/Json"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	*cmd.Application
	*errors.Errors
	Json.Json
}

func NewHandlers(app *cmd.Application) *handler {
	return &handler{
		Application: &cmd.Application{
			Config: app.Config,
			Logger: app.Logger,
			Models: app.Models,
		},
		Errors: errors.NewErrors(app.Logger),
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
