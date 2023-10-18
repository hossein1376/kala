package transfer

// HttpResponseError is meant to be used for documentation and should not be used for any other purposes
type httpResponseError struct {
	Message string `json:"message,omitempty" example:"error message"`
}

// HttpResponseSuccess is meant to be used for documentation and should not be used for any other purposes
type httpResponseSuccess struct {
	Data any `json:"data,omitempty" example:"an array or an object"`
}
