package response

import (
	"log/slog"
	"net/http"
)

type status string

const (
	OK    = "ok"
	Fail  = "fail"
	Error = "error"
)

type HttpResponse struct {
	Status     status `json:"status"`
	StatusCode uint   `json:"status_code,omitempty"`
	Message    any    `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Time       string `json:"time,omitempty"`
}

// Respond responses with the provided message and status code, it will return 500 if case of failure
func (res *Response) Respond(w http.ResponseWriter, r *http.Request, statusCode int, message any) {
	err := res.Write(w, statusCode, message, nil)
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
