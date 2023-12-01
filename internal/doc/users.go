package doc

import (
	"github.com/hossein1376/kala/internal/structure"
)

type createNewUserHandlerResponse struct {
	Data structure.User `json:"data,omitempty"`
}

type getUserByIDHandlerResponse struct {
	Data structure.User `json:"data,omitempty"`
}

type getAllUsersHandlerResponse struct {
	Data structure.GetAllUsersResponse `json:"data,omitempty"`
}
