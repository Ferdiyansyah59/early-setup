package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt"
)


func GenerateJWTToken(privateKey []byte, ttl time.Duration, userId int) (string, error){
	claims := BuildJWTClaims(ttl, userId)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
	
}