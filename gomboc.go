package main

import (
	"gomboc/api"
	"gomboc/environ"
	"time"

	"os"

	zerolog "github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func setZeroLogger() *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
	log.Logger = log.Logger.With().Timestamp().Caller().Logger()

	return &(log.Logger)
}

func main() {
	logger := setZeroLogger()             // set the global zloggger configuration
	lconf := environ.GombocConfigLoader() // load hub customer configuration

	logger.Info().Msg("Wellcome to Gomboc Suite")

	// Start HubAPI.
	//
	// Peer devices will register with the api. Control and management of devices will be done
	// through the administrator user interface.
	aconf := lconf.GombocServer.GombocServerAPI
	api.GombocAPI{Host: aconf.Host, Port: aconf.Port, AutoMigrate: false}.Initialize()
}
