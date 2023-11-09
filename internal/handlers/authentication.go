package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/password"
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
		h.Info(login, "status", transfer.StatusBadRequest, "error", err)
		h.BadRequestResponse(w, err)
		return
	}

	user, err := h.Models.User.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}) || errors.As(err, &transfer.ForbiddenError{}):
			h.Info(login, "status", transfer.StatusUnauthorized, "error", err, "ip", h.getIP(r))
			h.UnauthorizedResponse(w)
		default:
			h.Error(login, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	p := password.Password{
		Plaintext: &input.Password,
		Hash:      user.Password.Hash,
	}
	ok, err := p.ArgonMatches()
	if err != nil {
		h.Error(login, "status", transfer.StatusInternalServerError, "error", err)
		h.InternalServerErrorResponse(w)
		return
	}
	if !ok {
		h.Info(login, "status", transfer.StatusUnauthorized, "error", "password not match")
		h.UnauthorizedResponse(w)
		return
	}

	_, token, err := h.Config.JWT.Token.Encode(map[string]any{
		"user":   user,
		"expire": time.Now().Add(time.Hour * 24 * 7).Format(time.RFC3339),
	})
	if err != nil {
		h.Error(login, "status", transfer.StatusInternalServerError, "error", err)
		h.InternalServerErrorResponse(w)
		return
	}

	w.Header().Set("Bearer", token)
	resp := transfer.Resp{Data: structure.LoginResponse{Token: token}}
	h.Info(login, "status", transfer.StatusOK, "response", "login successful")
	h.OkResponse(w, resp)
}
