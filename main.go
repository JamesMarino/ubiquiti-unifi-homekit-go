package main

import (
	"github.com/jamesmarino/ubiquiti-unifi-homekit-go/pkg/homekit"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	configFilename := "config.yml"

	config, err := homekit.InitialiseConfiguration(configFilename)
	if err != nil {
		log.Fatal().Err(err).Msg("config file not loaded")
	}

	err = homekit.BroadcastDevices(config)
	if err != nil {
		log.Fatal().Err(err).Msg("can't broadcast devices")
	}
}
