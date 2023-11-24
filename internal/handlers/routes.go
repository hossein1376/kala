package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
)

const (
	login = "login handler"

	createNewUser  = "create new user handler"
	getAllUsers    = "get all users handler"
	getUserByID    = "get user by id handler"
	updateUserByID = "update user by id handler"
	deleteUserByID = "delete user by id handler"
)

func (h *handler) Router() *chi.Mux {
	// create new router
	r := chi.NewRouter()

	// standard middlewares
	r.Use(middleware.Recoverer)

	// custom middlewares
	r.Use(h.logger)

	r.Route("/api/v1", func(r chi.Router) {
		// public routes
		r.Group(func(r chi.Router) {
			// rate limiter middleware
			r.Use(httprate.LimitByRealIP(10, time.Minute))

			r.Post("/login", h.loginHandler)
		})

		// private routes
		r.Group(func(r chi.Router) {
			// JWT middlewares
			r.Use(jwtauth.Verifier(h.Config.JWT.Token), jwtauth.Authenticator(h.Config.JWT.Token))
			r.Use(h.checkJWT)

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
