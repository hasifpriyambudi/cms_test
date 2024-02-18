package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// CHeck Header
		if authHeader == "" {
			err := errors.New("unauthorized")
			panic(exceptions.NewAuthError(err))
		}

		// Check Contains Bearer
		if !strings.Contains(authHeader, "Bearer ") {
			err := errors.New("unauthorized")
			panic(exceptions.NewAuthError(err))
		}

		// Get Token & Validate
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := helpers.ValidateToken(authHeader)
		if err != nil {
			err = errors.New("unauthorized")
			panic(exceptions.NewAuthError(err))
		}

		// Cek Token Valid
		if !token.Valid {
			err := errors.New("unauthorized")
			panic(exceptions.NewAuthError(err))
		}

		// Set Context
		ctx.Set("token", authHeader)
		ctx.Next()
	}
}
