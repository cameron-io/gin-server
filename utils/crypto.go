package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func MatchPasswords(input_pw string, db_pw string) error {
	hashed_pw, err := HashPassword(input_pw)
	if err != nil {
		return errors.New("internal error")
	}
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed_pw),
		[]byte(db_pw))
}
