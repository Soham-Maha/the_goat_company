package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/vaibhavsijaria/TGC-be.git/models"
)

func GenerateJWT(user interface{}) (string, error) {
	var jwtKey = GetEnv("SECRET_KEY")
	expirationTime := time.Now().Add(24 * time.Hour)
	var email string

	switch u := user.(type) {
	case *models.Farmer:
		email = u.Email
	case *models.Investor:
		email = u.Email
	case *models.Vet:
		email = u.Email
	default:
		return "", errors.New("invalid user type")
	}

	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	var jwtKey = GetEnv("SECRET_KEY")

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
