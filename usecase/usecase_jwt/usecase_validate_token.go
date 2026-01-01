package usecasejwt

import (
	"early-setup/helper"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (usecase *usecaseJWTImpl) ValidaetJWTUsecase(signedToken string) (string ,error) {
	token ,err := jwt.Parse(signedToken, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
		}
		return signedToken, nil
	})
	if err != nil {
		return "0", err
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			err = fmt.Errorf("Validate: invalid")
			return "0", err
		} else {
			if userId, err := helper.Decrypt(string(claims["dat"].(string))); err != nil {
				return "0", err
			}else {
				return userId, nil
			}

		}
	}
}