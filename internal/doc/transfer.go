package doc

type httpResponseError struct {
	Message string `json:"message,omitempty" example:"error message"`
}

type httpResponseSuccess struct {
	Data any `json:"data,omitempty" example:"an array or an object"`
}
