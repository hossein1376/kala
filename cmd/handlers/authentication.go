package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/Errors"
	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/pkg/Password"
)

func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.LoginRequest
	err := h.ReadJSONiter(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.Models.User.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.As(err, &Errors.UserNotFound{}):
			h.UnauthorizedResponse(w, r)
		case errors.As(err, &Errors.UserDisabled{}):
			h.UnauthorizedResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	p := Password.Password{
		Plaintext: &input.Password,
		Hash:      string(user.Password),
	}
	ok, err := p.ArgonMatches()
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}
	if !ok {
		h.UnauthorizedResponse(w, r)
		return
	}

	_, token, err := h.Config.JWTToken.Encode(map[string]any{
		"user":   user,
		"expire": time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339),
	})
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	err = h.WriteJSONiter(w, http.StatusOK, map[string]any{"token": token}, nil)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}
}