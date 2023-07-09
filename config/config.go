package config

import (
	"github.com/rs/zerolog/log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT  string `envconfig:"PORT" default:"8000"`
	DBURL string `envconfig:"DB_URL" required:"true"`
}

var ENVConfig Config

func ReadConfig() error {
	err := envconfig.Process("bounty", &ENVConfig)
	if err != nil {
		log.Error().Err(err).Msg("error when reading configuration")
		return err
	}

	log.Info().Msg("success reading configuration")

	return nil
}
