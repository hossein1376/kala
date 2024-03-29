package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
)

func (h *handler) createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.Product
	err := h.ReadJson(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	err = h.Models.Product.Create(ctx, &input)
	if err != nil {
		h.InternalServerErrorResponse(w)
		return
	}

	h.CreatedResponse(w, input)
}

func (h *handler) getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, nil)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	product, err := h.Models.Product.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, err)
			return
		default:
			h.InternalServerErrorResponse(w)
			return
		}
	}

	h.OkResponse(w, product)
}

func (h *handler) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	products, err := h.Models.Product.GetAll(ctx)
	if err != nil {
		h.InternalServerErrorResponse(w)
		return
	}

	h.OkResponse(w, products)
}

func (h *handler) updateProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w)
		return
	}

	var input structure.ProductUpdate
	err := h.ReadJson(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	product, err := h.Models.Product.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, err)
		default:
			h.InternalServerErrorResponse(w)
		}
		return
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Quantity != nil {
		product.Quantity = *input.Quantity
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Status != nil {
		product.Status = *input.Status
	}
	if input.Available != nil {
		product.Available = *input.Available
	}
	if input.Rating != nil {
		product.Rating = *input.Rating
	}
	if input.RatingCount != nil {
		product.RatingCount = *input.RatingCount
	}

	err = h.Models.Product.UpdateByID(ctx, id, product)
	if err != nil {
		h.InternalServerErrorResponse(w)
		return
	}

	h.NoContentResponse(w)
}

func (h *handler) deleteProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := h.paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	err := h.Models.Product.DeleteByID(ctx, id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(w, err)
		default:
			h.InternalServerErrorResponse(w)
		}
		return
	}

	h.NoContentResponse(w)
}
