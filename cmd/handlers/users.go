package handlers

import (
	"net/http"

	"kala/internal/ent"
	"kala/internal/structure"
	"kala/pkg/Errors"
	"kala/pkg/Json"
	"kala/pkg/Password"
)

func (h *Handler) createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.UserRequest
	err := Json.ReadJSON(w, r, &input)
	if err != nil {
		Errors.BadRequestResponse(w, r, err)
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

	err = user.Password.Set(input.Password)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}

	// TODO: validations

	err = h.app.Models.User.CreateNewUser(user)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			Errors.BadRequestResponse(w, r, err)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = Json.WriteJSON(w, http.StatusCreated, user, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	user, err := h.app.Models.User.GetSingleUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
			return
		default:
			Errors.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	err = Json.WriteJSON(w, http.StatusOK, user, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.app.Models.User.GetAllUsers()
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
	err = Json.WriteJSON(w, http.StatusOK, users, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	var input structure.UserUpdateRequest
	err := Json.ReadJSON(w, r, &input)
	if err != nil {
		Errors.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.app.Models.User.GetSingleUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
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
		err = p.Set(*input.Password)
		if err != nil {
			Errors.InternalServerErrorResponse(w, r, err)
			return
		}
		user.Password = []byte(p.Hash)
	}

	err = h.app.Models.User.UpdateUserByID(id, user)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = Json.WriteJSON(w, http.StatusOK, user, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	err := h.app.Models.User.DeleteUserByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = Json.WriteJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}
