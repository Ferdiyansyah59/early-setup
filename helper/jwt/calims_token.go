package jwt

import (
	"early-setup/helper"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func BuildJWTClaims(ttl time.Duration, userId int) (claims jwt.MapClaims) {
	userIdString := strconv.Itoa(userId)
	encryptedUserId := helper.GenerateEncrypt(userIdString)
	now := time.Now().UTC()
	claims = make(jwt.MapClaims)
	claims["dat"] = encryptedUserId
	claims["exp"] = now.Add(ttl).Unix() 
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	return
}	