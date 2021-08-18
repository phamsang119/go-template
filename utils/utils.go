package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}

func IsDuplicateRecord(err error) bool {
	if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
		return true
	}
	return false
}
