package middleware

import (
	"lms_system/internal/domain/common"
	"lms_system/internal/domain/entity"
	"net/http"
	"strings"
)

func OnlyRoles(allowedRoles ...common.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxVal := r.Context().Value(common.UserContextKey)
			if ctxVal == nil {
				http.Error(w, "User context not found", http.StatusUnauthorized)
				return
			}

			userCtx, ok := ctxVal.(*entity.UserContext)
			if !ok {
				http.Error(w, "Invalid user context", http.StatusInternalServerError)
				return
			}

			userRole := strings.ToUpper(string(userCtx.Role))
			for _, role := range allowedRoles {
				if userRole == strings.ToUpper(string(role)) {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
