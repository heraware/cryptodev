package cmd

import (
	"fmt"
	"log"

	bitcoin "github.com/btcsuite/btcrpcclient"
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
			getNewAddress(bitcoinClient, &args)
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

func getNewAddress(bitcoinClient *bitcoin.Client, args *[]string) {
	account := "N/A"
	if len(*args) > 1 {
		account = (*args)[1]
	}
	address, err := bitcoinClient.GetNewAddress(account)
	if err != nil {
		log.Fatal(err)
	}
	result := fmt.Sprintf("Address: %s Account: %s", address, account)
	fmt.Println(result)
}
