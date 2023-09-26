package handlers

import (
	"net/http"
)

func (h *handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	h.Info("received request!")
	h.Respond(w, r, http.StatusOK, "Hello World")
}
