// middleware/auth_middleware.go
package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"golang_restaurant_management/utils"
)

// contextKey type prevents collisions in context keys
type contextKey string

const userKey contextKey = "userUUID"

// JWTAuth middleware checks for a valid JWT token
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// 2. Check format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// 3. Verify JWT
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// 4. Attach user UUID to request context
		ctx := context.WithValue(r.Context(), userKey, claims["uuid"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserUUID returns the user UUID stored in context by JWTAuth
func GetUserUUID(r *http.Request) (string, error) {
	if val, ok := r.Context().Value(userKey).(string); ok {
		return val, nil
	}
	return "", errors.New("user UUID not found in context")
}
