package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *handler) Router() *chi.Mux {
	// create new router
	r := chi.NewRouter()

	// standard middleware
	r.Use(middleware.Logger, middleware.Recoverer)

	// custom middlewares
	r.Use(h.logger)

	// routes
	r.Get("/", h.homeHandler)
	r.Route("/api/v1", func(r chi.Router) {
		// user routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.createNewUserHandler)
			r.Get("/", h.getAllUsersHandler)
			r.Get("/{id}", h.getUserByIDHandler)
			r.Patch("/{id}", h.updateUserByIDHandler)
			r.Delete("/{id}", h.deleteUserByIDHandler)
		})

		// product routes
		r.Route("/products", func(r chi.Router) {
			r.Post("/", h.createNewProductHandler)
			r.Get("/", h.getAllProductsHandler)
			r.Get("/{id}", h.getProductByIDHandler)
			r.Patch("/{id}", h.updateProductByIDHandler)
			r.Delete("/{id}", h.deleteProductByIDHandler)
		})

		// category routes
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", h.createNewCategoryHandler)
			r.Get("/", h.getAllCategoriesHandler)
			r.Get("/{id}", h.getCategoryByIDHandler)
		})

		// brand routes
		r.Route("/brands", func(r chi.Router) {
			r.Post("/", h.createNewBrandHandler)
			r.Get("/", h.getAllBrandsHandler)
			r.Get("/{id}", h.getBrandByIDHandler)
		})

		// comment routes
		r.Route("/comments", func(r chi.Router) {
			r.Post("/{id}", h.createNewProductCommentHandler)
			r.Get("/{id}", h.getAllProductCommentsHandler)
			r.Get("/{id}/{commentID}", h.getProductCommentByIDHandler)
		})

		// seller routes
		r.Route("/sellers", func(r chi.Router) {
			r.Post("/", h.createNewSellerHandler)
			r.Get("/", h.getAllSellersHandler)
			r.Get("/{id}", h.getSellerByIDHandler)
		})

		// address routes
		r.Route("/addresses", func(r chi.Router) {
			r.Post("/", h.createNewAddressHandler)
			r.Get("/", h.getAllAddressesHandler)
			r.Get("/{id}", h.getAddressByIDHandler)
		})

		// image route
		r.Route("/images", func(r chi.Router) {
			r.Post("/", h.createNewImageHandler)
		})
	})

	return r
}
