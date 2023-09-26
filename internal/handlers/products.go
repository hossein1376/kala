package handlers

import (
	"net/http"

	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/structure"
)

func (h *handler) createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.Product
	err := h.Read(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	product, err := h.Models.Product.CreateNewProduct(input)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.BadRequestResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.Respond(w, r, http.StatusCreated, product)
}

func (h *handler) getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	product, err := h.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
			return
		default:
			h.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	h.Respond(w, r, http.StatusOK, product)
}

func (h *handler) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.Models.Product.GetAllProducts()
	if err != nil {
		h.InternalServerErrorResponse(w, r, err)
		return
	}

	h.Respond(w, r, http.StatusOK, products)
}

func (h *handler) updateProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	var input structure.ProductUpdate
	err := h.Read(w, r, &input)
	if err != nil {
		h.BadRequestResponse(w, r, err)
		return
	}

	product, err := h.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
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
	if input.Images != nil {
		product.Edges.Image = input.Images
	}
	if input.Category != nil {
		product.Edges.Category = input.Category
	}
	if input.Brand != nil {
		product.Edges.Brand = input.Brand
	}

	err = h.Models.Product.UpdateProductByID(product, id)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.BadRequestResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.Respond(w, r, http.StatusNoContent, product)
}

func (h *handler) deleteProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.NotFoundResponse(w, r, nil)
		return
	}

	err := h.Models.Product.DeleteProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.NotFoundResponse(w, r, err)
		default:
			h.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	h.Respond(w, r, http.StatusOK, nil)
}
