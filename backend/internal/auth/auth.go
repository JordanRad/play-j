package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(passwordString string) (string, error) {
	password := []byte(passwordString)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashString string, passwordString string) bool {
	password := []byte(passwordString)
	hash := []byte(hashString)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}

func GenerateJWT() string {
	return "Hello"
}

func ValidateJWT(token string) string{
	return "Hello"
}

