package endpoints

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

type contextKey string

const emailContextKey contextKey = "email"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "request does not contains an authorization header"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOAK_URL"))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "erro to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailgo"})
		// verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true}) //pula a verificação do clienteID do keycloak
		_, err = verifier.Verify(r.Context(), tokenString)

		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "invalid token"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]

		ctx := context.WithValue(r.Context(), emailContextKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
