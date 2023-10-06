package handlers

import (
	"github.com/hossein1376/kala/config"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/Json"
)

// JsonHandler is alias for default json handler package
type JsonHandler = Json.Jsoniter

type handler struct {
	*config.Application
	*transfer.Response
	JsonHandler
}

func NewHandlers(app *config.Application) *handler {
	return &handler{
		Application: &config.Application{
			Config: app.Config,
			Logger: app.Logger,
			Models: app.Models,
			RDB:    app.RDB,
		},
		Response: transfer.NewResponse(app.Logger, JsonHandler{}),
	}
}
