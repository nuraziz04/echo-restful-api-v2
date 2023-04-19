package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
