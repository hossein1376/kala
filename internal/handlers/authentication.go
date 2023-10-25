package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/Password"
)

// loginHandler godoc
//
// @Summary            Login
// @Description        Sign-in into the application
// @Tags               authentication
// @Accept             json
// @Produce            json
// @Param              request body      structure.LoginRequest    true     "user credentials"
// @Success            200     {object}  doc.loginHandlerResponse           "single object containing JWT token"
// @Failure            400     {object}  doc.httpResponseError              "bad input"
// @Failure            401     {object}  doc.httpResponseError              "not authorized"
// @Failure            500     {object}  doc.httpResponseError              "unexpected error"
// @Router             /login [post]
func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.LoginRequest
	err := h.ReadJson(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.Models.User.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}) || errors.As(err, &transfer.ForbiddenError{}):
			h.UnauthorizedResponse(w, r)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	p := Password.Password{
		Plaintext: &input.Password,
		Hash:      user.Password.Hash,
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

	_, token, err := h.Config.JWT.Token.Encode(map[string]any{
		"user":   user,
		"expire": time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339),
	})
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Bearer", token)
	resp := transfer.HttpResponse{Data: structure.LoginResponse{Token: token}}
	h.StatusOKResponse(w, r, resp)
}
