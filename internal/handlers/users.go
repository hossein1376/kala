package handlers

import (
	"errors"
	"net/http"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/Password"
)

func (h *handler) createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.UserCreationRequest
	err := h.ReadJson(w, r, &input)
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
		case errors.As(err, &transfer.BadRequestError{}):
			h.BadRequestResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.StatusCreatedResponse(w, r, transfer.HttpResponse{Data: u})
}

func (h *handler) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	user, err := h.Models.User.GetByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.StatusOKResponse(w, r, transfer.HttpResponse{Data: user})
}

func (h *handler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Models.User.GetAll()
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	h.StatusOKResponse(w, r, transfer.HttpResponse{Data: users})
}

func (h *handler) updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	var input structure.UserUpdateRequest
	err := h.ReadJson(w, r, &input)
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
		case errors.As(err, &transfer.NotFoundError{}):
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
		user.Password.Hash = p.Hash
	}

	err = h.Models.User.UpdateByID(id, user)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.StatusOKResponse(w, r, transfer.HttpResponse{Data: user})
}

func (h *handler) deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	err := h.Models.User.DeleteByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.StatusNoContentResponse(w, r)
}
