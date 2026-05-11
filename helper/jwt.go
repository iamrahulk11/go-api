package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(claims map[string]interface{}, Audience, Issuer, Secret string, ExpiresInMinute int) (string, error) {
	tokenClaims := jwt.MapClaims{
		"iss": Issuer,
		"aud": Audience,
		"exp": time.Now().Add(time.Minute * time.Duration(ExpiresInMinute)).Unix(),
		"iat": time.Now().Unix(),
	}

	// add custom claims
	for k, v := range claims {
		tokenClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString([]byte(Secret))
}

// ValidateToken parses and validates JWT
func ValidateToken(tokenString, Audience, Issuer, Secret string, ExpiresInMinute int) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithAudience(Audience),
		jwt.WithIssuer(Issuer),
		jwt.WithLeeway(time.Duration(ExpiresInMinute)),
	)

	claims := jwt.MapClaims{}
	token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
