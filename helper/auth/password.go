package auth

import (
	"grpctest/utils"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	utils.ErrorPanic(err)

	return string(hashedPassword)
}

func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
