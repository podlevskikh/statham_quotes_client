package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type Specs struct {
	HTTPPost string `envconfig:"HTTP_PORT" default:"8080"`
}

var specs Specs

func init() {
	err := godotenv.Overload()
	if err != nil {
		fmt.Printf("failed to load .env file %s\n", err)
	}
	err = envconfig.Process("", &specs)
	if err != nil {
		zerolog.DefaultContextLogger.Fatal().Err(err).Msg("missing envs")
	}
}
