package cmd

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new node - Example: cryptodev create bitcoin",
	Run: func(cmd *cobra.Command, args []string) {
		docker := clients.NewDockerClient()
		docker.CreateNode("bitcoin")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
