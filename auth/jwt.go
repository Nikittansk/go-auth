package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(username, email string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}
