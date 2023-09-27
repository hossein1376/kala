package response

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/hossein1376/kala/pkg/Json"
)

const (
	OK    = "ok"
	Fail  = "fail"
	Error = "error"
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

type HttpResponse struct {
	Status     string `json:"status"`
	StatusCode uint   `json:"status_code,omitempty"`
	Message    any    `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Time       string `json:"time,omitempty"`
}

// Respond responses with the provided message and status code, it will return 500 if case of failure
func (res *Response) Respond(w http.ResponseWriter, r *http.Request, statusCode int, message any) {
	err := res.WriteJson(w, statusCode, message, nil)
	if err != nil {
		res.logInternalError(r, err)
		w.WriteHeader(500)
	}
}

// logInternalError logs the error details to the standard logger
func (res *Response) logInternalError(req *http.Request, err error) {
	res.Error("internal error",
		slog.String("error_message", err.Error()),
		slog.String("request_method", req.Method),
		slog.String("request_url", req.URL.String()))
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (res *Response) InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	res.logInternalError(r, err)

	response := HttpResponse{
		Status:     Error,
		StatusCode: http.StatusInternalServerError,
		Message:    http.StatusText(http.StatusNotFound),
		Time:       time.Now().Format(time.RFC3339),
	}
	res.Respond(w, r, http.StatusInternalServerError, response)
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (res *Response) NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := http.StatusText(http.StatusNotFound)
	if err != nil {
		msg = err.Error()
	}

	response := HttpResponse{
		Status:     Fail,
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Time:       time.Now().Format(time.RFC3339),
	}
	res.Respond(w, r, http.StatusNotFound, response)
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (res *Response) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	response := HttpResponse{
		Status:     Error,
		StatusCode: http.StatusMethodNotAllowed,
		Message:    fmt.Sprintf("the %s method is not supported for this content", r.Method),
		Time:       time.Now().Format(time.RFC3339),
	}
	res.Respond(w, r, http.StatusMethodNotAllowed, response)
}

// BadRequestResponse indicates that the request is not what the server expects
func (res *Response) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	msg := http.StatusText(http.StatusBadRequest)
	if err != nil {
		msg = err.Error()
	}

	response := HttpResponse{
		Status:     Error,
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Time:       time.Now().Format(time.RFC3339),
	}
	res.Respond(w, r, http.StatusBadRequest, response)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (res *Response) DuplicateRequestResponse(w http.ResponseWriter, r *http.Request) {
	response := HttpResponse{
		Status:     Fail,
		StatusCode: http.StatusNotAcceptable,
		Message:    "the content you're trying to add already exists",
		Time:       time.Now().Format(time.RFC3339),
	}
	res.Respond(w, r, http.StatusNotAcceptable, response)
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Response) ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	res.Respond(w, r, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

// UnauthorizedResponse responds when the user's action is not allowed
func (res *Response) UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	res.Respond(w, r, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}
