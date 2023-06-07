package handlers

import (
	"net/http"
)

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	h.app.Logger.Info("received request!")
	w.Write([]byte("Hello World"))
}
