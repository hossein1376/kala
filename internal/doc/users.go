package doc

import (
	"github.com/hossein1376/kala/internal/structure"
)

type getUserByIDHandlerResponse struct {
	Data structure.User `json:"data,omitempty"`
}

type createNewUserHandlerResponse struct {
	Data structure.User `json:"data,omitempty"`
}
