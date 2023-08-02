package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/structure"
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
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}

	// TODO
	image := &structure.Image{
		Name:    header.Filename,
		Path:    path,
		Caption: "",
		Width:   0,
		Height:  0,
		Size:    float64(header.Size),
		Owner:   0,
	}
	err = h.app.Models.Image.Insert(image)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r, err)
		case ent.IsConstraintError(err):
			h.error.BadRequestResponse(w, r, err)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}
}
