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
	Data []structure.User `json:"data,omitempty"`
}

type updateUserByIDHandlerResponse struct {
	Data structure.User `json:"data,omitempty"`
}
