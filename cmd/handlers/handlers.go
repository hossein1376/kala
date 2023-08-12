package handlers

import (
	"github.com/hossein1376/kala/cmd"
	"github.com/hossein1376/kala/internal/response"
	"github.com/hossein1376/kala/pkg/Json"
)

// JsonHandler is alias for default json handler package
type JsonHandler = Json.Jsoniter

type handler struct {
	*cmd.Application
	*response.Response
	JsonHandler
}

func NewHandlers(app *cmd.Application) *handler {
	return &handler{
		Application: &cmd.Application{
			Config: app.Config,
			Logger: app.Logger,
			Models: app.Models,
		},
		Response: response.NewResponse(app.Logger, JsonHandler{}),
	}
}
