package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (h *handler) createNewImageHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		h.error.BadRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	err = os.MkdirAll("/etc/www/images", os.ModePerm)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}

	path := fmt.Sprintf("/etc/www/images/%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))
	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// make record in db

}
