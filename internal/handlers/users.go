package handlers

import (
	"errors"
	"net/http"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/password"
	"github.com/hossein1376/kala/pkg/validator"
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
		h.Info(createNewUser, "status", transfer.StatusBadRequest, "error", err)
		h.BadRequestResponse(w, err)
		return
	}

	user, v := new(structure.User), validator.New()

	user.Username = input.Username
	v.Check("username",
		validator.Case{Cond: validator.MinLength(user.Username, 4), Msg: "must be at least 4 characters"},
		validator.Case{Cond: validator.MaxLength(user.Username, 20), Msg: "must be at most 20 characters"},
	)

	v.Check("password",
		validator.Case{Cond: validator.RangeLength(input.Password, 4, 50), Msg: "must be between 4 and 50 characters"},
	)

	if input.FirstName != nil {
		user.FirstName = input.FirstName
		v.Check("first name",
			validator.Case{Cond: validator.NotEmpty(*user.FirstName), Msg: "must not be empty"},
		)
	}
	if input.LastName != nil {
		user.LastName = input.LastName
		v.Check("lastname",
			validator.Case{Cond: validator.NotEmpty(*user.LastName), Msg: "must not be empty"},
		)
	}
	if input.Email != nil {
		user.Email = input.Email
		v.Check("email",
			validator.Case{Cond: validator.Matches(*user.Email, validator.EmailRX), Msg: "must be valid"},
		)
	}
	if input.Phone != nil {
		user.Phone = input.Phone
		v.Check("phone",
			validator.Case{Cond: validator.IsNumber(*user.Phone), Msg: "must be only numbers"},
			validator.Case{Cond: validator.RangeLength(*user.Phone, 2, 20), Msg: "must be between 2 to 20 digits"},
		)
	}

	if ok := v.Valid(); !ok {
		response := transfer.Resp{Message: v.Errors}
		h.Info(createNewUser, "status", transfer.StatusBadRequest, "error", v.Errors)
		h.Response(w, http.StatusBadRequest, response)
		return
	}

	err = user.Password.ArgonSet(input.Password)
	if err != nil {
		h.Error(createNewUser, "status", transfer.StatusInternalServerError, "error", err)
		h.InternalServerErrorResponse(w)
		return
	}

	err = h.Models.User.Create(user)
	if err != nil {
		switch {
		case errors.As(err, &transfer.BadRequestError{}):
			h.Info(createNewUser, "status", transfer.StatusBadRequest, "error", err)
			h.BadRequestResponse(w, err)
		default:
			h.Error(createNewUser, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	response := transfer.Resp{Data: user}
	h.Info(createNewUser, "status", transfer.StatusCreated, "response", response)
	h.CreatedResponse(w, response)
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
		h.Info(getUserByID, "status", transfer.StatusNotFound, "error", "invalid param")
		h.NotFoundResponse(w)
		return
	}

	user, err := h.Models.User.GetByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(getUserByID, "status", transfer.StatusNotFound, "error", err)
			h.NotFoundResponse(w)
		default:
			h.Error(getUserByID, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	response := transfer.Resp{Data: user}
	h.Info(getUserByID, "status", transfer.StatusOK, "response", response)
	h.OkResponse(w, response)
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
		h.Info(getAllUsers, "status", transfer.StatusBadRequest, "error", err)
		h.BadRequestResponse(w, err)
		return
	}

	if query.Page < 1 || query.Count < 1 {
		h.Info(getAllUsers, "status", transfer.StatusNotFound, "error", "invalid page or count")
		h.BadRequestResponse(w)
		return
	}

	users, count, err := h.Models.User.GetAll(query)
	if err != nil {
		h.Error(getAllUsers, "status", transfer.StatusInternalServerError, "error", err)
		h.InternalServerErrorResponse(w)
		return
	}

	pages := count / query.Count
	if count%query.Count != 0 {
		pages++
	}
	resp := structure.GetAllUsersResponse{
		Users:       users,
		CurrentPage: query.Page,
		TotalPages:  pages,
		PerPage:     query.Count,
	}

	response := transfer.Resp{Data: resp}
	h.Info(getAllUsers, "status", transfer.StatusOK, "response", response)
	h.OkResponse(w, response)
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
// @Success            204
// @Failure            400     {object}  doc.httpResponseError                "bad input"
// @Failure            404     {object}  doc.httpResponseError                "user not found"
// @Failure            500     {object}  doc.httpResponseError                "unexpected error"
// @Router             /users/{id} [patch]
func (h *handler) updateUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.Info(updateUserByID, "status", transfer.StatusNotFound, "error", "invalid param")
		h.NotFoundResponse(w)
		return
	}

	var input structure.UserUpdateRequest
	err := h.ReadJson(w, r, &input)
	if err != nil {
		h.Info(updateUserByID, "status", transfer.StatusBadRequest, "error", err)
		h.BadRequestResponse(w, err)
		return
	}

	if input == (structure.UserUpdateRequest{}) {
		err = errors.New("empty request")
		h.Info(updateUserByID, "status", transfer.StatusBadRequest, "error", err)
		h.BadRequestResponse(w, err)
		return
	}

	user, err := h.Models.User.GetByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(updateUserByID, "status", transfer.StatusNotFound, "error", err)
			h.NotFoundResponse(w, err)
		default:
			h.Error(updateUserByID, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	if input.Password != nil {
		var p password.Password
		err = p.ArgonSet(*input.Password)
		if err != nil {
			h.Error(updateUserByID, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
			return
		}
		user.Password.Hash = p.Hash
	}

	if input.UserName != nil {
		user.Username = *input.UserName
	}
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Phone = input.Phone

	err = h.Models.User.UpdateByID(id, user)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(updateUserByID, "status", transfer.StatusNotFound, "error", err)
			h.NotFoundResponse(w, err)
		default:
			h.Error(updateUserByID, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	h.Info(updateUserByID, "status", transfer.StatusNoContent)
	h.NoContentResponse(w)
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
		h.Info(deleteUserByID, "status", transfer.StatusNotFound, "error", "invalid param")
		h.NotFoundResponse(w)
		return
	}

	err := h.Models.User.DeleteByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(deleteUserByID, "status", transfer.StatusNotFound, "error", err)
			h.NotFoundResponse(w, err)
		default:
			h.Error(deleteUserByID, "status", transfer.StatusInternalServerError, "error", err)
			h.InternalServerErrorResponse(w)
		}
		return
	}

	h.Info(deleteUserByID, "status", transfer.StatusNoContent, "error", err)
	h.NoContentResponse(w)
}
