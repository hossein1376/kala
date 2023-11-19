package transfer

import (
	"log/slog"
	"net/http"

	"github.com/hossein1376/kala/pkg/Json"
)

type Responder struct {
	*slog.Logger
	Json.Json
}

func NewResponse(logger *slog.Logger, json Json.Json) *Responder {
	return &Responder{
		Logger: logger,
		Json:   json,
	}
}

type Resp struct {
	Message any `json:"message,omitempty"`
	Data    any `json:"data,omitempty"`
}

// Response is a general function which responses with the provided message and status code,
// it will return 500 if case of failure
func (res *Responder) Response(w http.ResponseWriter, statusCode int, message any) {
	err := res.WriteJson(w, statusCode, message, nil)
	if err != nil {
		res.Error("writing json response", "error", err, "data", message)
		w.WriteHeader(500)
	}
}

// OkResponse means everything went as expected
func (res *Responder) OkResponse(w http.ResponseWriter, data any) {
	res.Response(w, http.StatusOK, data)
}

// CreatedResponse indicates that requested resource(s) have been successfully created
func (res *Responder) CreatedResponse(w http.ResponseWriter, data any) {
	res.Response(w, http.StatusCreated, data)
}

// NoContentResponse means the operation was successful, and server has nothing more to say about it
func (res *Responder) NoContentResponse(w http.ResponseWriter) {
	res.Response(w, http.StatusNoContent, nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (res *Responder) BadRequestResponse(w http.ResponseWriter, err ...error) {
	msg := http.StatusText(http.StatusBadRequest)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := Resp{Message: msg}
	res.Response(w, http.StatusBadRequest, response)
}

// UnauthorizedResponse responds when user is not authorized
func (res *Responder) UnauthorizedResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusUnauthorized)}
	res.Response(w, http.StatusUnauthorized, response)
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Responder) ForbiddenResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusForbidden)}
	res.Response(w, http.StatusForbidden, response)
}

// NotFoundResponse will return with classic 404 error message.
// if error message is provided, it will return that instead.
func (res *Responder) NotFoundResponse(w http.ResponseWriter, err ...error) {
	msg := http.StatusText(http.StatusNotFound)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := Resp{Message: msg}
	res.Response(w, http.StatusNotFound, response)
}

// DuplicateRequestResponse indicates a similar request has already been handled.
func (res *Responder) DuplicateRequestResponse(w http.ResponseWriter) {
	response := Resp{Message: "the content you're trying to add already exists"}
	res.Response(w, http.StatusNotAcceptable, response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly.
func (res *Responder) InternalServerErrorResponse(w http.ResponseWriter) {
	response := Resp{Message: http.StatusText(http.StatusInternalServerError)}
	res.Response(w, http.StatusInternalServerError, response)
}
