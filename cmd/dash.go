package cmd

import (
	"log"

	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// dashCmd represents the bitcoin command
var dashCmd = &cobra.Command{
	Use:   "dash",
	Short: "Run RPC command over Dash Node (getinfo, getnewaddress, listaccounts)",
	Run: func(cmd *cobra.Command, args []string) {
		bitcoinClient := clients.NewBitcoinClient(27001)

		if len(args) == 0 {
			log.Fatalln(`Action not provided.
Actions list:
 - getinfo
 - getnewaddress <ACCOUNT>
 - listaccounts
 - newblocks <AMOUNT OF NEW BLOCKS>
 - send <ADDRESS> <AMOUNT OF BTC>`)
		}
		runAction(&args, bitcoinClient)
	},
}

func init() {
	RootCmd.AddCommand(dashCmd)
}
