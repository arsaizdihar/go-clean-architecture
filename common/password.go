package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPassword(hashPassword string, plainPassword string) (bool, error) {
	hashPW := []byte(hashPassword)
	if err := bcrypt.CompareHashAndPassword(hashPW, []byte(plainPassword)); err != nil {
		return false, err
	}
	return true, nil
}
