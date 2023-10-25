package handlers

import (
	"errors"
	"net/http"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/Password"
)

// createNewUserHandler godoc
//
// @Summary            Create user
// @Description        creates a new user and returns the newly created data
// @Tags               user management
// @Accept             json
// @Produce            json
// @Param              request body      structure.UserCreationRequest  true  "user data"
// @Success            201     {object}  doc.createNewUserHandlerResponse     "single object containing user's data"
// @Failure            400     {object}  doc.httpResponseError                "bad input"
// @Failure            500     {object}  doc.httpResponseError                "unexpected error"
// @Router             /users [post]
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

// getUserByIDHandler godoc
//
// @Summary          Get user
// @Description      get a single user by id
// @Tags             user management
// @Accept           json
// @Produce          json
// @Param            id   path      int           true               "user ID"
// @Success          200  {object}  doc.getUserByIDHandlerResponse   "single object containing user's data"
// @Failure          400  {object}  doc.httpResponseError            "invalid ID"
// @Failure          404  {object}  doc.httpResponseError            "user not found"
// @Failure          500  {object}  doc.httpResponseError            "unexpected error"
// @Router           /users/{id} [get]
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
			h.NotFoundResponse(w, r, nil)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.StatusOKResponse(w, r, transfer.HttpResponse{Data: user})
}

// getAllUsersHandler godoc
//
// @Summary          Get all users
// @Description      get all users
// @Tags             user management
// @Accept           json
// @Produce          json
// @Param            page   query     int            true              "page number"
// @Param            count  query     int           true               "number of items per page"
// @Success          200    {object}  doc.getAllUsersHandlerResponse   "array of user objects"
// @Failure          500    {object}  doc.httpResponseError            "unexpected error"
// @Router           /users [get]
func (h *handler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := new(structure.GetAllUsersRequest)
	if err := h.decoder.Decode(query, r.URL.Query()); err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	if query.Page < 1 || query.Count < 1 {
		h.BadRequestResponse(w, r, nil)
		return
	}

	users, count, err := h.Models.User.GetAll(query)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	pages := count / query.Count
	if count%query.Count != 0 {
		pages++
	}
	response := structure.GetAllUsersResponse{
		Users:       users,
		CurrentPage: query.Page,
		TotalPages:  pages,
		PerPage:     query.Count,
	}

	h.StatusOKResponse(w, r, transfer.HttpResponse{Data: response})
}

// createNewUserHandler godoc
//
// @Summary            Update user
// @Description        update a user data by its ID
// @Tags               user management
// @Accept             json
// @Produce            json
// @Param              request body      structure.UserUpdateRequest  true    "user data"
// @Param              id      path      int                          true    "user ID"
// @Success            200     {object}  doc.updateUserByIDHandlerResponse    "single object containing user's data"
// @Failure            400     {object}  doc.httpResponseError                "bad input"
// @Failure            404     {object}  doc.httpResponseError                "user not found"
// @Failure            500     {object}  doc.httpResponseError                "unexpected error"
// @Router             /users/{id} [patch]
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

// deleteUserByIDHandler godoc
//
// @Summary            Delete user
// @Description        delete a user by its ID
// @Tags               user management
// @Accept             json
// @Produce            json
// @Param              id      path      int              true    "User ID"
// @Success            204
// @Failure            404     {object}  doc.httpResponseError   "user not found"
// @Failure            500     {object}  doc.httpResponseError   "unexpected error"
// @Router             /users/{id} [delete]
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
