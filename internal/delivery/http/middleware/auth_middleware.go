package middleware

import (
	"context"
	"errors"
	"lms_system/internal/domain/common"
	"lms_system/internal/domain/entity"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		userCtx, err := validateToken(token)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), common.UserContextKey, userCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(tokenString string) (*entity.UserContext, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username, ok := claims["preferred_username"].(string)
	if !ok || username == "" {
		return nil, errors.New("missing username in token")
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return nil, errors.New("missing sub in token")
	}

	role := common.RoleUser

	if ra, ok := claims["realm_access"].(map[string]interface{}); ok {
		if rawRoles, ok := ra["roles"].([]interface{}); ok {
			for _, r := range rawRoles {
				roleStr, _ := r.(string)
				switch strings.ToUpper(roleStr) {
				case "ROLE_ADMIN":
					role = common.RoleAdmin
				case "ROLE_TEACHER":
					role = common.RoleTeacher
				}
			}
		}
	}

	// For testing: give admin role to admin user
	// This is a temporary hack for development/testing
	if username == "admin" {
		role = common.RoleAdmin
	}

	return &entity.UserContext{
		UserID:   sub,
		Username: username,
		Role:     common.UserRole(role),
	}, nil
}
