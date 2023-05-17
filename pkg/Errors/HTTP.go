package Errors

import (
	"fmt"
	"net/http"

	"kala/pkg/Json"
	"kala/pkg/Logger"

	"golang.org/x/exp/slog"
)

// logError logs the error details to the standard logger
func logError(r *http.Request, err error) {
	Logger.Logger.Error("internal error",
		slog.String("error_message", err.Error()),
		slog.String("request_method", r.Method),
		slog.String("request_url", r.URL.String()))
}

// errorResponse responses with the error message
func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := Json.WriteJSON(w, status, message, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(500)
	}
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	errorResponse(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// NotFoundResponse returns the classic 404
func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this content", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// BadRequestResponse indicates that the request is not what the server expects
func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// FailedValidationResponse indicates that the data validation resulted with error
func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	errorResponse(w, r, http.StatusNotAcceptable, errors)
}

// EditConflictResponse indicates the rare instances of edit race
func EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	errorResponse(w, r, http.StatusConflict, message)
}

// RateLimitExceededResponse kicks in after rateLimiter middleware
func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
}

// InvalidCredentialsResponse indicates the invalid input credentials
func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

// MismatchedInfoResponse indicates that the request body info and the route related info does not match
func MismatchedInfoResponse(w http.ResponseWriter, r *http.Request) {
	message := "id in the route and id in the request body does not match"
	errorResponse(w, r, http.StatusNotAcceptable, message)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func DuplicateRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "the content you're trying to add already exists"
	errorResponse(w, r, http.StatusNotAcceptable, message)
}

// ForbiddenResponse indicates that the action is not allowed
func ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

// UnauthorizedResponse responds when the user's action is not allowed
func UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, r, http.StatusForbidden, http.StatusText(http.StatusUnauthorized))
}
