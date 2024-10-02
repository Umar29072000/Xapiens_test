package middleware

import (
	"net/http"
	"strings"
	"Xapiens_test/services"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusForbidden)
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		_, err := services.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}