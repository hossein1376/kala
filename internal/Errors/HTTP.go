// Package Errors provides common error responses.
// Note that it relies on Json package.
package Errors

import (
	"fmt"
	"net/http"

	"github.com/hossein1376/kala/pkg/Json"

	"golang.org/x/exp/slog"
)

type Errors struct {
	*slog.Logger
	Json.Json
}

func NewErrors(logger *slog.Logger) *Errors {
	json := Json.NewJson()
	return &Errors{
		Logger: logger,
		Json:   json,
	}
}

// logInternalError logs the error details to the standard logger
func (e *Errors) logInternalError(r *http.Request, err error) {
	e.Error("internal error",
		slog.String("error_message", err.Error()),
		slog.String("request_method", r.Method),
		slog.String("request_url", r.URL.String()))
}

// errorResponse responses with the error message
func (e *Errors) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := e.WriteJSON(w, status, message, nil)
	if err != nil {
		e.logInternalError(r, err)
		w.WriteHeader(500)
	}
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (e *Errors) InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.logInternalError(r, err)
	e.errorResponse(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (e *Errors) NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		e.errorResponse(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		e.errorResponse(w, r, http.StatusNotFound, err.Error())
	}
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (e *Errors) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this content", r.Method)
	e.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// BadRequestResponse indicates that the request is not what the server expects
func (e *Errors) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	e.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// FailedValidationResponse indicates that the data validation resulted with error
func (e *Errors) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	e.errorResponse(w, r, http.StatusNotAcceptable, errors)
}

// EditConflictResponse indicates the rare instances of edit race
func (e *Errors) EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	e.errorResponse(w, r, http.StatusConflict, message)
}

// RateLimitExceededResponse kicks in after rateLimiter middleware
func (e *Errors) RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	e.errorResponse(w, r, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
}

// InvalidCredentialsResponse indicates the invalid input credentials
func (e *Errors) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	e.errorResponse(w, r, http.StatusUnauthorized, message)
}

// MismatchedInfoResponse indicates that the request body info and the route related info does not match
func (e *Errors) MismatchedInfoResponse(w http.ResponseWriter, r *http.Request) {
	message := "id in the route and id in the request body does not match"
	e.errorResponse(w, r, http.StatusNotAcceptable, message)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (e *Errors) DuplicateRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "the content you're trying to add already exists"
	e.errorResponse(w, r, http.StatusNotAcceptable, message)
}

// ForbiddenResponse indicates that the action is not allowed
func (e *Errors) ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	e.errorResponse(w, r, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

// UnauthorizedResponse responds when the user's action is not allowed
func (e *Errors) UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	e.errorResponse(w, r, http.StatusForbidden, http.StatusText(http.StatusUnauthorized))
}
