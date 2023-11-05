package transfer

import (
	"fmt"
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

// Responder responses with the provided message and status code, it will return 500 if case of failure
func (res *Response) Responder(w http.ResponseWriter, r *http.Request, statusCode int, message any) {
	err := res.WriteJson(w, statusCode, message, nil)
	if err != nil {
		res.logInternalError(err)
		w.WriteHeader(500)
	}
}

// logInternalError logs the error details to the standard logger
func (res *Response) logInternalError(err error) {
	res.Error("writing json response", "status", InternalServerError, "error", err)
}

// OkResponse means everything went as expected
func (res *Response) OkResponse(w http.ResponseWriter, r *http.Request, data any) {
	res.Responder(w, r, http.StatusOK, data)
}

// CreatedResponse indicates that requested resource(s) have been successfully created
func (res *Response) CreatedResponse(w http.ResponseWriter, r *http.Request, data any) {
	res.Responder(w, r, http.StatusCreated, data)
}

// NoContentResponse means the operation was successful, and server has nothing more to say about it
func (res *Response) NoContentResponse(w http.ResponseWriter, r *http.Request) {
	res.Responder(w, r, http.StatusNoContent, nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (res *Response) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := http.StatusText(http.StatusBadRequest)
	if err != nil {
		msg = err.Error()
	}

	response := Resp{Message: msg}
	res.Responder(w, r, http.StatusBadRequest, response)
}

// UnauthorizedResponse responds when user is not authorized
func (res *Response) UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	res.Responder(w, r, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Response) ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	res.Responder(w, r, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (res *Response) NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := http.StatusText(http.StatusNotFound)
	if err != nil {
		msg = err.Error()
	}

	response := Resp{Message: msg}
	res.Responder(w, r, http.StatusNotFound, response)
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (res *Response) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	response := Resp{Message: fmt.Sprintf("the %s method is not supported for this content", r.Method)}
	res.Responder(w, r, http.StatusMethodNotAllowed, response)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (res *Response) DuplicateRequestResponse(w http.ResponseWriter, r *http.Request) {
	response := Resp{Message: "the content you're trying to add already exists"}
	res.Responder(w, r, http.StatusNotAcceptable, response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (res *Response) InternalServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	response := Resp{Message: http.StatusText(http.StatusInternalServerError)}
	res.Responder(w, r, http.StatusInternalServerError, response)
}
