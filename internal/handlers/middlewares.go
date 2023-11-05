package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"

	"github.com/hossein1376/kala/internal/transfer"
)

const (
	logger   = "logger"
	checkJWT = "check JWT"
)

func (h *handler) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Info(logger, "method", r.Method, "url", r.URL.String(), "ip", h.getIP(r))

		defer func() {

		}()

		next.ServeHTTP(w, r)
	})
}

func (h *handler) checkJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			h.Error(checkJWT, "status", transfer.InternalServerError, "error", err)
			h.InternalServerErrorResponse(w, r)
			return
		}
		expireString, ok := claims["expire"].(string)
		if !ok {
			h.Info(checkJWT, "status", transfer.Unauthorized, "error", "bad expiration type")
			h.UnauthorizedResponse(w, r)
			return
		}
		expire, err := time.Parse(time.RFC3339, expireString)
		if err != nil {
			h.Info(checkJWT, "status", transfer.Unauthorized, "error", "failed to parse expire date")
			h.UnauthorizedResponse(w, r)
			return
		}

		if time.Now().After(expire) {
			h.Info(checkJWT, "status", transfer.Unauthorized, "error", "expired token")
			h.UnauthorizedResponse(w, r)
			return
		}

		user, ok := claims["user"].(map[string]any)
		if !ok {
			h.Info(checkJWT, "status", transfer.Unauthorized, "error", "bad user type")
			h.UnauthorizedResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
