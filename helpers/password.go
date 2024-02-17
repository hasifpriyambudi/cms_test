package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicError(err)

	return string(passwordHash)
}

func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
