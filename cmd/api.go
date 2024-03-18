package cmd

import (
	"gomboc/api"

	"github.com/spf13/cobra"
)

var host string
var port int
var migrate bool

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start Gomboc API",
	Long:  "Start the Gomboc API with a host (default 0.0.0.0) and a port (default 7020)",
	Run: func(cmd *cobra.Command, args []string) {
		// Start Gomboc API.
		//
		// Peer devices will register with the api. Control and management of devices will be
		// done through the administrator user interface.
		api.GombocAPI{Host: host, Port: port, AutoMigrate: migrate}.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.Flags().StringVarP(&host, "host", "", "0.0.0.0", "API host address")
	apiCmd.Flags().IntVarP(&port, "port", "", 7020, "API port number")
	apiCmd.Flags().BoolVarP(&migrate, "migrate", "", false, "Auto migrate models")
}
