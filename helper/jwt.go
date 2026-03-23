package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SecretKey       string
	Issuer          string
	Audience        string
	ExpiresInMinute int
}

func (j *JWT) CreateToken(claims map[string]interface{}) (string, error) {
	tokenClaims := jwt.MapClaims{
		"iss": j.Issuer,
		"aud": j.Audience,
		"exp": time.Now().Add(time.Minute * time.Duration(j.ExpiresInMinute)).Unix(),
		"iat": time.Now().Unix(),
	}

	// add custom claims
	for k, v := range claims {
		tokenClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString([]byte(j.SecretKey))
}

// ValidateToken parses and validates JWT
func (j *JWT) ValidateToken(tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithAudience(j.Audience),
		jwt.WithIssuer(j.Issuer),
		jwt.WithLeeway(time.Duration(j.ExpiresInMinute)),
	)

	claims := jwt.MapClaims{}
	token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
