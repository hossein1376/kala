// Package est: Error State Transfer
package est

import (
	"net/http"
)

type Error struct {
	Err            error
	HTTPStatusCode int
}

func (e Error) Error() string {
	if e.Err == nil {
		return http.StatusText(e.HTTPStatusCode)
	}
	return e.Err.Error()
}

func BadRequestError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusBadRequest}
}

func UnauthorizedError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusUnauthorized}
}

func ForbiddenError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusForbidden}
}

func NotFoundError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusNotFound}
}

func NotAcceptableError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusNotAcceptable}
}

func ConflictError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusConflict}
}

func TooManyRequestsError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusTooManyRequests}
}

func InternalError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusInternalServerError}
}

func BadGatewayError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusBadGateway}
}

func GatewayTimeoutError(err error) Error {
	return Error{Err: err, HTTPStatusCode: http.StatusGatewayTimeout}
}
