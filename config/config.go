package config

import (
	"early-setup/helper"
	"early-setup/model/dto"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var Config dto.ServerConfig

func init() {
	err := loadConfig()
	helper.PanicIfError(err)
}


func loadConfig() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("Cannot find .env file")
	}

	err = env.Parse(&Config)
	return err
}