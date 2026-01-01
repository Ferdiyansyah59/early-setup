package usecasejwt

import (
	"early-setup/config"
	"early-setup/helper/jwt"
	"early-setup/model/dto/response"
	"time"
)

func (usecase *usecaseJWTImpl) GenerateJWTUsecase(userId int) (res response.TokenResponseDTO, err error) {
	expiredTime := config.Config.JWT.TokenExp
	secretKey := config.Config.JWT.SecretKey
	token, err := jwt.GenerateJWTToken(secretKey, expiredTime, userId)

	layout := "2006-01-02 15:04:05 MST"
	expiredAt := time.Now().Add(expiredTime).Format(layout)
	
	if err == nil {
		res.Token = token
		res.ExpiredAt = expiredAt
	}
	return
}