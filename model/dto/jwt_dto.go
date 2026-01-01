package dto

import "time"

type JWTconfigDto struct {
	SecretKey []byte `env:"JWT_SECRET_KEY"`
	TokenExp time.Duration `env:"JWT_ACCESS_TOKEN_TTL"`
}