package cmd

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:.`,
	Run: func(cmd *cobra.Command, args []string) {
		docker := clients.NewDockerClient()
		docker.CreateNode("bitcoin")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
