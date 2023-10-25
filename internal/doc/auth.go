package doc

import (
	"github.com/hossein1376/kala/internal/structure"
)

type loginHandlerResponse struct {
	Data structure.LoginResponse `json:"data,omitempty"`
}
