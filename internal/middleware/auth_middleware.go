package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pranavkolte/chat-server-websocket/internal/util"
)

var authRoute string = "/api/v1/auth"

var exculdedPaths = map[string]struct{}{
	authRoute + "/login":  {},
	authRoute + "/signup": {},
}

func JWTAuthMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if _, ok := exculdedPaths[r.URL.Path]; ok {
			nextHandler.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or Invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return util.GetJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Acess Token Invalid", http.StatusUnauthorized)
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}
