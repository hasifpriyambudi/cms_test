package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func AuthJWT(auth domain.AuthDatabase) (string, error) {
	expTime := time.Now().Add(time.Minute * 60)
	claims := &entity.JwtClaims{
		ID:       auth.Id,
		Name:     auth.Name,
		Username: auth.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Kebenaran",
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(privateKey)
}

func parseToken(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
	}

	return []byte(privateKey), nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, parseToken)
}
