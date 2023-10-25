package handlers

import (
	"github.com/gorilla/schema"

	"github.com/hossein1376/kala/config"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/Json"
)

type handler struct {
	*config.Application
	*transfer.Response
	decoder *schema.Decoder
}

func NewHandlers(app *config.Application) *handler {
	return &handler{
		Application: &config.Application{
			Config: app.Config,
			Logger: app.Logger,
			Models: app.Models,
			RDB:    app.RDB,
		},
		Response: transfer.NewResponse(app.Logger, Json.Json{}),
		decoder:  schema.NewDecoder(),
	}
}
