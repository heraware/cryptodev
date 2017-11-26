package cmd

import (
	"fmt"
	"log"

	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// bitcoinCashCmd represents the bitcoin command
var bitcoinCashCmd = &cobra.Command{
	Use:   "bitcoin-cash",
	Short: "Run RPC command over Bitcoin Cash Node (getinfo, getnewaddress, listaccounts)",
	Run: func(cmd *cobra.Command, args []string) {
		bitcoinClient := clients.NewBitcoinClient()

		if len(args) == 0 {
			log.Fatalln(`Action not provided.
Actions list:
 - getinfo
 - getnewaddress <ACCOUNT>
 - listaccounts
 - newblocks <AMOUNT OF NEW BLOCKS>
 - send <ADDRESS> <AMOUNT OF BTC>`)
		}
		switch args[0] {
		case "getinfo":
			getInfo(bitcoinClient)
		case "getnewaddress":
			getNewAddress(bitcoinClient, &args)
		case "listaccounts":
			listAccounts(bitcoinClient)
		case "newblocks":
			newBlocks(bitcoinClient, &args)
		case "send":
			send(bitcoinClient, &args)
		default:
			log.Fatal(fmt.Sprintf("Action %s is not valid.", args[0]))
		}
	},
}

func init() {
	RootCmd.AddCommand(bitcoinCashCmd)
}
