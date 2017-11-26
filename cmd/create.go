package cmd

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new node - Example: cryptodev create bitcoin",
	Long: `
Available nodes:
- bitcoin
- bitcoin-cash
- bitcoin-gold
- zcash
- zen
- litecoin
- ethereum
- ethereum-classic`,
	Run: func(cmd *cobra.Command, args []string) {
		docker := clients.NewDockerClient()
		if len(args) > 0 {
			nodeName := args[0]
			docker.CreateNode(nodeName)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
