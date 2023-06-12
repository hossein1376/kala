package handlers

import (
	"net/http"

	"kala/internal/ent"
	"kala/internal/structure"
	"kala/pkg/Password"
)

func (h *Handler) createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.UserRequest
	err := h.json.ReadJSONiter(w, r, &input)
	if err != nil {
		h.error.BadRequestResponse(w, r, err)
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
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}

	// TODO: validations

	err = h.app.Models.User.CreateNewUser(user)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.error.BadRequestResponse(w, r, err)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusCreated, user, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	user, err := h.app.Models.User.GetSingleUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
			return
		default:
			h.error.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	err = h.json.WriteJSONiter(w, http.StatusOK, user, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.app.Models.User.GetAllUsers()
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
	err = h.json.WriteJSONiter(w, http.StatusOK, users, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	var input structure.UserUpdateRequest
	err := h.json.ReadJSONiter(w, r, &input)
	if err != nil {
		h.error.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.app.Models.User.GetSingleUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
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
			h.error.InternalServerErrorResponse(w, r, err)
			return
		}
		user.Password = []byte(p.Hash)
	}

	err = h.app.Models.User.UpdateUserByID(id, user)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusOK, user, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	err := h.app.Models.User.DeleteUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusNoContent, nil, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}
