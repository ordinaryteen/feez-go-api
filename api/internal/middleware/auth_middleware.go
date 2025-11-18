package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type contextKey string

const userContextKey = contextKey("userID")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Validasi "Tanda Tangan" token
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID := claims.Subject

		ctx := context.WithValue(r.Context(), userContextKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext adalah 'helper' biar 'handler' gampang ngambil ID
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userContextKey).(string)
	return userID, ok
}
