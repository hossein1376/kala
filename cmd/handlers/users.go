package handlers

import (
	"net/http"

	"kala/internal/Errors"
	"kala/internal/Json"
	"kala/internal/ent"
)

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {

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
