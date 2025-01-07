package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsPasswordMatched(password string) error {
	hash, hash_err := HashPassword(password)
	if hash_err != nil {
		return errors.New("internal error")
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
