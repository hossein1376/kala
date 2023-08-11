package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/response"
	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/pkg/Password"
)

func (h *handler) createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.UserRequest
	err := h.Read(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	var user structure.User
	user.Username = input.UserName
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}

	err = user.Password.ArgonSet(input.Password)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	// TODO: validations

	u, err := h.Models.User.Create(user)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.BadRequestResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	resp := response.HttpResponse{
		Status:     response.OK,
		StatusCode: http.StatusCreated,
		Data:       u,
		Time:       time.Now().Format(time.RFC3339),
	}
	h.Respond(w, r, http.StatusCreated, resp)
}

func (h *handler) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	user, err := h.Models.User.GetByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
			return
		default:
			h.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	resp := response.HttpResponse{
		Status:     response.OK,
		StatusCode: http.StatusOK,
		Data:       user,
		Time:       time.Now().Format(time.RFC3339),
	}
	h.Respond(w, r, http.StatusOK, resp)
}

func (h *handler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Models.User.GetAll()
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}
	resp := response.HttpResponse{
		Status:     response.OK,
		StatusCode: http.StatusOK,
		Data:       users,
		Time:       time.Now().Format(time.RFC3339),
	}
	h.Respond(w, r, http.StatusOK, resp)
}

func (h *handler) updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	var input structure.UserUpdateRequest
	err := h.Read(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	if input == (structure.UserUpdateRequest{}) {
		h.BadRequestResponse(w, r, errors.New("empty request"))
		return
	}

	user, err := h.Models.User.GetByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	if input.UserName != nil {
		user.Username = *input.UserName
	}
	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.Password != nil {
		var p Password.Password
		err = p.ArgonSet(*input.Password)
		if err != nil {
			h.InternalServerErrorResponse(w, r, err)
			return
		}
		user.Password = []byte(p.Hash)
	}

	err = h.Models.User.UpdateByID(id, user)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	resp := response.HttpResponse{
		Status:     response.OK,
		StatusCode: http.StatusOK,
		Data:       user,
		Time:       time.Now().Format(time.RFC3339),
	}
	h.Respond(w, r, http.StatusOK, resp)
}

func (h *handler) deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	err := h.Models.User.DeleteByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.Respond(w, r, http.StatusNoContent, nil)
}
