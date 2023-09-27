package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
)

func (h *handler) Router() *chi.Mux {
	// create new router
	r := chi.NewRouter()

	// standard middlewares
	r.Use(middleware.Logger, middleware.Recoverer)

	// custom middlewares
	r.Use(h.logger)

	// public routes
	r.Group(func(r chi.Router) {
		// rate limiter middleware
		r.Use(httprate.LimitByRealIP(10, time.Minute))

		r.Post("/login", h.loginHandler)
	})

	// private routes
	r.Group(func(r chi.Router) {
		// JWT middlewares
		r.Use(jwtauth.Verifier(h.Config.JWTToken))
		r.Use(jwtauth.Authenticator, h.checkJWT)

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
		})
	})

	return r
}
