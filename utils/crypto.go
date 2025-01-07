package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func MatchPasswords(input_pw string, db_pw string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(db_pw),
		[]byte(input_pw))
}
