package cmd

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an existing node - Example: cryptodev run bitcoin",
	Run: func(cmd *cobra.Command, args []string) {
		docker := clients.NewDockerClient()
		docker.RunNode("bitcoin")
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
