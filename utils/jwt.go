package utils

import (
	"XM_assignment/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(username string, jwtKey string) (string, error) {
	claims := &domain.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
