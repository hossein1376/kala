package handlers

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info("received request!")
	w.Write([]byte("Hello World"))
}
