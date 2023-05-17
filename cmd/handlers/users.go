package handlers

import (
	"net/http"

	"kala/internal/ent"
	"kala/internal/structure"
	"kala/pkg/Errors"
	"kala/pkg/Json"
)

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
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

	err = app.Models.User.CreateNewUser(user)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r)
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	user, err := app.Models.User.GetSingleUserByID(id)
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

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := app.Models.User.GetAllUsers()
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

func updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {

}
