package handlers

import (
	"net/http"

	"kala/internal/ent"
	"kala/internal/structure"
	"kala/pkg/Errors"
	"kala/pkg/Json"
)

func createNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var input structure.Product
	err := Json.ReadJSON(w, r, &input)
	if err != nil {
		Errors.BadRequestResponse(w, r, err)
		return
	}

	product, err := app.Models.Product.CreateNewProduct(input)
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			Errors.BadRequestResponse(w, r, err)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = Json.WriteJSON(w, http.StatusCreated, product, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r)
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	product, err := app.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
			return
		default:
			Errors.InternalServerErrorResponse(w, r, err)
			return
		}
	}

	err = Json.WriteJSON(w, http.StatusOK, product, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.Models.Product.GetAllProducts()
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}

	err = Json.WriteJSON(w, http.StatusOK, products, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}
}

func updateProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r)
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	var input structure.ProductUpdate
	err := Json.ReadJSON(w, r, &input)
	if err != nil {
		Errors.BadRequestResponse(w, r, err)
		return
	}

	product, err := app.Models.Product.GetSingleProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
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
	if input.SubCategory != nil {
		product.Edges.SubCategory = input.SubCategory
	}
	if input.Brand != nil {
		product.Edges.Brand = input.Brand
	}

	err = app.Models.Product.UpdateProductByID(product, id)
}

func deleteProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := paramInt(r)
	if id == 0 {
		Errors.NotFoundResponse(w, r)
		return
	}

	err := app.Models.Product.DeleteProductByID(id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			Errors.NotFoundResponse(w, r)
		default:
			Errors.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = Json.WriteJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		Errors.InternalServerErrorResponse(w, r, err)
		return
	}

}
