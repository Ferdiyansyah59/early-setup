package usecasejwt

import "early-setup/model/dto/response"

type UsecaseJWT interface {
	GenerateJWTUsecase(userId int) (res response.TokenResponseDTO, err error)
	ValidaetJWTUsecase(signedToken string) (string ,error)
}