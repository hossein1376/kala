package transfer

import (
	"log/slog"
	"net/http"

	"github.com/hossein1376/kala/pkg/Json"
)

type Response struct {
	*slog.Logger
	Json.Json
}

func NewResponse(logger *slog.Logger, json Json.Json) *Response {
	return &Response{
		Logger: logger,
		Json:   json,
	}
}

type Resp struct {
	Message any `json:"message,omitempty"`
	Data    any `json:"data,omitempty"`
}

// Responder is a generic function which responses with the provided message and status code,
// it will return 500 if case of failure
func (res *Response) Responder(w http.ResponseWriter, statusCode int, message any) {
	err := res.WriteJson(w, statusCode, message, nil)
	if err != nil {
		res.logInternalError(err, message)
		w.WriteHeader(500)
	}
}

// OkResponse means everything went as expected
func (res *Response) OkResponse(w http.ResponseWriter, data any) {
	res.Responder(w, http.StatusOK, data)
}

// CreatedResponse indicates that requested resource(s) have been successfully created
func (res *Response) CreatedResponse(w http.ResponseWriter, data any) {
	res.Responder(w, http.StatusCreated, data)
}

// NoContentResponse means the operation was successful, and server has nothing more to say about it
func (res *Response) NoContentResponse(w http.ResponseWriter) {
	res.Responder(w, http.StatusNoContent, nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (res *Response) BadRequestResponse(w http.ResponseWriter, err ...error) {
	msg := http.StatusText(http.StatusBadRequest)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := Resp{Message: msg}
	res.Responder(w, http.StatusBadRequest, response)
}

// UnauthorizedResponse responds when user is not authorized
func (res *Response) UnauthorizedResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusUnauthorized)}
	res.Responder(w, http.StatusUnauthorized, response)
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Response) ForbiddenResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusForbidden)}
	res.Responder(w, http.StatusForbidden, response)
}

// NotFoundResponse will return with classic 404 error message.
// if error message is provided, it will return that instead.
func (res *Response) NotFoundResponse(w http.ResponseWriter, err ...error) {
	msg := http.StatusText(http.StatusNotFound)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := Resp{Message: msg}
	res.Responder(w, http.StatusNotFound, response)
}

// DuplicateRequestResponse indicates a similar request has already been handled.
func (res *Response) DuplicateRequestResponse(w http.ResponseWriter) {
	response := Resp{Message: "the content you're trying to add already exists"}
	res.Responder(w, http.StatusNotAcceptable, response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly.
func (res *Response) InternalServerErrorResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusInternalServerError)}
	res.Responder(w, http.StatusInternalServerError, response)
}

// logInternalError logs the error when writing json response fails
func (res *Response) logInternalError(err error, data any) {
	res.Error("writing json response", "error", err, "data", data)
}
