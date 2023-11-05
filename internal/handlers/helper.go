package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// paramInt extracts ID as an integer value from the URL parameter.
func (h *handler) paramInt(r *http.Request, name string) int {
	id, err := strconv.Atoi(chi.URLParam(r, name))
	if err != nil {
		return 0
	}
	return id
}

// background runs a task in a new goroutine. It also recovers from panics
func (h *handler) background(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				h.Error("background goroutine panic, failed to recover", "error", err)
			}
		}()

		fn()
	}()
}

func (h *handler) getIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
