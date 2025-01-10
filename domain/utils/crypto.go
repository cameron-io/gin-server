package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func MatchPasswords(inputPW string, dbPW string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(dbPW),
		[]byte(inputPW))
}
