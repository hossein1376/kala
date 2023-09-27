package transfer

import (
	"net/http"
)

type BadRequestError struct {
	Err error
}

func (e BadRequestError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusBadRequest)
	}
	return e.Err.Error()
}

type UnauthorizedError struct {
	Err error
}

func (e UnauthorizedError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusUnauthorized)
	}
	return e.Err.Error()
}

type ForbiddenError struct {
	Err error
}

func (e ForbiddenError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusForbidden)
	}
	return e.Err.Error()
}

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusNotFound)
	}
	return e.Err.Error()
}

type NotAcceptableError struct {
	Err error
}

func (e NotAcceptableError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusNotAcceptable)
	}
	return e.Err.Error()
}

type ConflictError struct {
	Err error
}

func (e ConflictError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusConflict)
	}
	return e.Err.Error()
}

type TooManyRequestsError struct {
	Err error
}

func (e TooManyRequestsError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusTooManyRequests)
	}
	return e.Err.Error()
}

type InternalError struct {
	Err error
}

func (e InternalError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusInternalServerError)
	}
	return e.Err.Error()
}

type BadGatewayError struct {
	Err error
}

func (e BadGatewayError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusBadGateway)
	}
	return e.Err.Error()
}

type GatewayTimeoutError struct {
	Err error
}

func (e GatewayTimeoutError) Error() string {
	if e.Err == nil {
		return http.StatusText(http.StatusGatewayTimeout)
	}
	return e.Err.Error()
}
