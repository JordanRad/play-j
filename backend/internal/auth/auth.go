package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenClaims struct {
	AccountID uint
	Email     string
	Exp       float64
}

// JWT secret used to create the signature
var secret = []byte("FHvdhfgvdJ98332")

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

func GenerateJWT(accountID uint, email string) (string, error) {
	// Set custom claims to JWT
	claims := &jwt.MapClaims{
		"iss": "playj-account-service",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"data": map[string]string{
			"accountID": strconv.Itoa(int(accountID)),
			"email":     email,
		},
	}

	// Sign the token and add the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", fmt.Errorf("error generating JWT: %w", err)
	}

	return tokenString, nil
}

func ExtractJWTCLaims(tokenString string) (TokenClaims, error) {
	//Parse the JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return TokenClaims{}, fmt.Errorf("error casting claims: %w", err)
	}

	// Extract the token's claims
	claims := token.Claims.(jwt.MapClaims)

	// Parse the claims to a struct
	exp := claims["exp"].(float64)
	data := claims["data"].(map[string]interface{})
	accountID := data["accountID"].(string)
	email := data["email"].(string)

	accID, err := strconv.Atoi(accountID)

	if err != nil {
		return TokenClaims{}, fmt.Errorf("error converting the accountID to uint: %w", err)
	}
	tokenClaims := TokenClaims{
		AccountID: uint(accID),
		Email:     email,
		Exp:       exp,
	}

	return tokenClaims, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	// //Parse the JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, fmt.Errorf("error signing the token: %w", err)
		}
	}
	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}
	return true, nil
}
