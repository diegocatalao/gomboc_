package cmd

import (
	"os"
  
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomboc",
	Short: "Start Gomboc application service",
	Long:  "Start Gomboc application services: api, rdv or client",
}

func setDefaultZeroLogger() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
	log.Logger = log.Logger.With().Timestamp().Caller().Logger()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Execute() {
	setDefaultZeroLogger() // set the global zloggger configuration

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
