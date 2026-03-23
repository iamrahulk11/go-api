package middlewares

import (
	"context"
	"net/http"
	"strings"

	"user-mapping/helper"
)

type JWTMiddleware struct {
	JWT *helper.JWT
}

func (m *JWTMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.JWT.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = WithClaims(ctx, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type claimsKeyType string

const claimsKey claimsKeyType = "jwtClaims"

func WithClaims(ctx context.Context, claims map[string]interface{}) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func GetClaims(ctx context.Context) (map[string]interface{}, bool) {
	claims, ok := ctx.Value(claimsKey).(map[string]interface{})
	return claims, ok
}
