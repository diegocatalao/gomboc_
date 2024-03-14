package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nodeRDVHost string
var nodeRDVPort int

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start Gomboc Node service (the client service)",
	Long:  "Start Gomboc Node service with a RDV hub host (rdv-ip) and a port (rdv-port 7021)",
	Run: func(cmd *cobra.Command, args []string) {
		if nodeRDVHost == "" {
			panic("RDV host cannot to be empty!")
		}
		fmt.Println("rdv called")
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringVarP(&nodeRDVHost, "rdv-host", "", "", "RDV host address")
	nodeCmd.Flags().IntVarP(&nodeRDVPort, "rdv-port", "", 7020, "RDV port number")
}
