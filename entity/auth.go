package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	ID       int
	Name     string
	Username string
	jwt.RegisteredClaims
}
