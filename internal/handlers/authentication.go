package handlers

import (
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/response"
	"github.com/hossein1376/kala/internal/structure"
	est "github.com/hossein1376/kala/pkg/EST"
	"github.com/hossein1376/kala/pkg/Password"
)

func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.LoginRequest
	err := h.Read(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.Models.User.GetByUsername(input.Username)
	if err != nil {
		h.Respond(w, r, err.(est.Error).HTTPStatusCode, err.Error())
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

	w.Header().Set("Bearer", token)
	resp := response.HttpResponse{
		Status:     response.OK,
		StatusCode: http.StatusOK,
		Data:       map[string]any{"token": token},
		Time:       time.Now().Format(time.RFC3339),
	}
	h.Respond(w, r, http.StatusOK, resp)
}
