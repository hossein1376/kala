package handlers

import (
	"net/http"

	"kala/internal/ent"
	"kala/internal/structure"
)

func (h *handler) createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.Product
	err := h.json.ReadJSONiter(w, r, &input)
	if err != nil {
		h.error.BadRequestResponse(w, r, err)
		return
	}

	product, err := h.app.Models.Product.CreateNewProduct(input)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.error.BadRequestResponse(w, r, err)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusCreated, product, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *handler) getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	product, err := h.app.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
			return
		default:
			h.error.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	err = h.json.WriteJSONiter(w, http.StatusOK, product, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *handler) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.app.Models.Product.GetAllProducts()
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusOK, products, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *handler) updateProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	var input structure.ProductUpdate
	err := h.json.ReadJSONiter(w, r, &input)
	if err != nil {
		h.error.BadRequestResponse(w, r, err)
		return
	}

	product, err := h.app.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
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
	//if input.Images != nil {
	//	product.Edges.Image = input.Images
	//}
	if input.Category != nil {
		product.Edges.Category = input.Category
	}
	if input.SubCategory != nil {
		product.Edges.SubCategory = input.SubCategory
	}
	if input.Brand != nil {
		product.Edges.Brand = input.Brand
	}

	err = h.app.Models.Product.UpdateProductByID(product, id)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			h.error.BadRequestResponse(w, r, err)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusNoContent, product, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *handler) deleteProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r, "id")
	if id == 0 {
		h.error.NotFoundResponse(w, r)
		return
	}

	err := h.app.Models.Product.DeleteProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			h.error.NotFoundResponse(w, r)
		default:
			h.error.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.json.WriteJSONiter(w, http.StatusOK, nil, nil)
	if err != nil {
		h.error.InternalServerErrorResponse(w, r, err)
		return
	}

}
