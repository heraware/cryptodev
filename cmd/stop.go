package cmd

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an existing node - Example: cryptodev stop bitcoin",
	Run: func(cmd *cobra.Command, args []string) {
		docker := clients.NewDockerClient()
		docker.StopNode("bitcoin")
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)
}
