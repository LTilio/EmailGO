package endpoints

import (
	"EmailGO/internal/infra/credential"
	"context"
	"net/http"

	"github.com/go-chi/render"
)

type contextKey string

const emailContextKey contextKey = "email"

type ValidateTokenFunc func(token string, ctx context.Context) (string, error)

var ValidateToken ValidateTokenFunc = credential.ValidateToken

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "request does not contains an authorization header"})
			return
		}

		email, err := ValidateToken(tokenString, r.Context())

		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), emailContextKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
