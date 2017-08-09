package cmd

import (
	"fmt"
	"log"

	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// bitcoinCmd represents the bitcoin command
var bitcoinCmd = &cobra.Command{
	Use:   "bitcoin",
	Short: "Run RPC command over Bitcoin Node (getinfo, getnewaddress, listaccounts)",
	Run: func(cmd *cobra.Command, args []string) {
		bitcoinClient := clients.NewBitcoinClient()

		if len(args) == 0 {
			log.Fatalln("Action not provided. Actions list: getinfo, getnewaddress, listaccounts")
		}
		switch args[0] {
		case "getinfo":
			info, err := bitcoinClient.GetInfo()
			log.Fatal(err)
			fmt.Println(info)
		case "getnewaddress":
			account := ""
			if len(args) > 1 {
				account = args[1]
			}
			address, _ := bitcoinClient.GetNewAddress(account)
			fmt.Println(address)
		case "listaccounts":
			accounts, err := bitcoinClient.ListAccounts()
			log.Fatal(err)
			fmt.Println(accounts)
		default:
			log.Fatal(fmt.Sprintf("Action %s is not valid.", args[0]))
		}
	},
}

func init() {
	RootCmd.AddCommand(bitcoinCmd)
}
