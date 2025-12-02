package middleware

import (
	"net/http"
	"strings"

	"github.com/qobilovvv/test_tasks/auth/pkg/helpers"
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			helpers.ResponseError(w, http.StatusUnauthorized, "missing Authorization header")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			helpers.ResponseError(w, http.StatusUnauthorized, "invalid Authorization header format")
			return
		}

		token := parts[1]
		_, userType, err := helpers.DecodeAccessToken(token)
		if err != nil {
			helpers.ResponseError(w, http.StatusUnauthorized, err.Error())
			return
		}

		path := r.URL.Path
		protected := map[string]bool{
			"/roles":           true,
			"/roles/{id}":      true,
			"/roles/create":    true,
			"/sysusers/create": true,
		}

		if protected[path] && userType != "sysuser" {
			helpers.ResponseError(w, http.StatusForbidden, "permission denied")
			return
		}

		next.ServeHTTP(w, r)
	})
}
