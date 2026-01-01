package dto

import "time"

type JWTconfigDto struct {
	secretKey []byte `env:"JWT_SECRET_KEY"`
	tokenExp time.Duration `env:"JWT_ACCESS_TOKEN_TTL"`
}