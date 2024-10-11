package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string) bool {
	hash, hash_err := HashPassword(password)
	if hash_err != nil {
		return false
	}
	compare_err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return compare_err == nil
}
