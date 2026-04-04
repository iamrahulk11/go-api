package middlewares

import (
	"strings"

	"user-mapping/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(jwtHelper *helper.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtHelper.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		// store claims in gin context
		c.Set("jwtClaims", claims)

		c.Next()
	}
}

func GetClaims(c *gin.Context) (map[string]interface{}, bool) {
	claims, exists := c.Get("jwtClaims")
	if !exists {
		return nil, false
	}

	data, ok := claims.(map[string]interface{})
	return data, ok
}
