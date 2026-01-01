package usecasejwt

import (
	"early-setup/model/dto"

	"github.com/go-playground/validator/v10"
)

type usecaseJWTImpl struct {
	validate *validator.Validate
	secretKey dto.JWTconfigDto
}

func NewUsecaseJWT(validate *validator.Validate, secretKey dto.JWTconfigDto) UsecaseJWT {
	return &usecaseJWTImpl{
		validate: validate,
		secretKey: secretKey,
	}
}