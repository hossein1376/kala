package handlers

import (
	"net/http"
)

func (h *handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	h.Info("received request!")

	err := h.WriteJSON(w, http.StatusOK, "Hello World", nil)
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}
}
