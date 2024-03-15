package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rdvHost string
var rdvPort int

var rdvCmd = &cobra.Command{
	Use:   "rdv",
	Short: "Start Gomboc RDV service (rendezvous server)",
	Long:  "Start Gomboc RDV service with a host (default 0.0.0.0) and a port (default 7021)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rdv called")
	},
}

func init() {
	rootCmd.AddCommand(rdvCmd)

	rdvCmd.Flags().StringVarP(&rdvHost, "host", "", "", "RDV host address")
	rdvCmd.Flags().IntVarP(&rdvPort, "port", "", 7020, "RDV port number")
}
