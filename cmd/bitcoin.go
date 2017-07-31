package cmd

import (
	"fmt"

	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// bitcoinCmd represents the bitcoin command
var bitcoinCmd = &cobra.Command{
	Use:   "bitcoin",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		bitcoinClient := clients.NewBitcoinClient()

		switch args[0] {
		case "getinfo":
			info, _ := bitcoinClient.GetInfo()
			fmt.Println(info)
		case "getnewaddress":
			account := ""
			if len(args) > 1 {
				account = args[1]
			}
			address, _ := bitcoinClient.GetNewAddress(account)
			fmt.Println(address)
		case "listaccounts":
			accounts, _ := bitcoinClient.ListAccounts()
			fmt.Println(accounts)
		}
	},
}

func init() {
	RootCmd.AddCommand(bitcoinCmd)
}
