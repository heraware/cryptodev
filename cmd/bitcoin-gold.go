package cmd

import (
	"log"

	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// bitcoinGoldCmd represents the bitcoin command
var bitcoinGoldCmd = &cobra.Command{
	Use:   "bitcoin-gold",
	Short: "Run RPC command over Bitcoin Gold Node (getinfo, getnewaddress, listaccounts)",
	Run: func(cmd *cobra.Command, args []string) {
		bitcoinClient := clients.NewBitcoinClient(23001)

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
	RootCmd.AddCommand(bitcoinGoldCmd)
}
