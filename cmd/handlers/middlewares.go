package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func (h *handler) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func (h *handler) checkJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			h.InternalServerErrorResponse(w, r, err)
			return
		}
		expireString, ok := claims["expire"].(string)
		if !ok {
			h.Info("bad expiration type")
			h.UnauthorizedResponse(w, r)
			return
		}
		expire, err := time.Parse(time.RFC3339, expireString)
		if err != nil {
			h.Info("failed to parse expire date")
			h.UnauthorizedResponse(w, r)
			return
		}

		if time.Now().After(expire) {
			h.Info("expired token")
			h.UnauthorizedResponse(w, r)
			return
		}

		user, ok := claims["user"].(map[string]any)
		if !ok {
			h.Info("bad user type")
			h.UnauthorizedResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
