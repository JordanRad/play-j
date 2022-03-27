package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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

func GenerateJWT(accountID uint) (string, error) {
	// Create the JWT key used to create the signature
	var secret = []byte("FHvdhfgvdJ98332")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(accountID)),
		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
	})
	token, err := claims.SignedString(secret)

	if err != nil {
		return "", fmt.Errorf("error generating JWT: %w", err)
	}

	return token, nil
}

func ValidateJWT(token string) string {
	return "Hello"
}
