package handlers

import (
	"kala/cmd"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var app cmd.Application

func NewHandler(application cmd.Application) *chi.Mux {
	// initialize the application state instance for the handlers package
	app = application

	// create new router
	r := chi.NewRouter()

	// standard middleware
	r.Use(middleware.Logger, middleware.Recoverer)

	// custom middlewares

	// routes
	r.Get("/", homeHandler)
	r.Route("/api/v1", func(r chi.Router) {
		// user routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", createNewUserHandler)
			r.Get("/", getAllUsersHandler)
			r.Get("/{id}", getUserByIDHandler)
			r.Patch("/{id}", updateUserByIDHandler)
			r.Delete("/{id}", deleteUserByIDHandler)
		})

		// product routes
		r.Route("/products", func(r chi.Router) {
			r.Post("/", createNewProductHandler)
			r.Get("/", getAllProductsHandler)
			r.Get("/{id}", getProductByIDHandler)
			r.Patch("/{id}", updateProductByIDHandler)
			r.Delete("/{id}", deleteProductByIDHandler)
		})

	})

	return r
}
