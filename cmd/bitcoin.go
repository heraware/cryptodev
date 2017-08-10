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
			getInfo(bitcoinClient)
		case "getnewaddress":
			getNewAddress(bitcoinClient, &args)
		case "listaccounts":
			listAccounts(bitcoinClient)
		default:
			log.Fatal(fmt.Sprintf("Action %s is not valid.", args[0]))
		}
	},
}

func init() {
	RootCmd.AddCommand(bitcoinCmd)
}

func getNewAddress(bitcoinClient *bitcoin.Client, args *[]string) {
	account := ""
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

func listAccounts(bitcoinClient *bitcoin.Client) {
	accounts, err := bitcoinClient.ListAccounts()
	if err != nil {
		log.Fatal(err)
	}
	for accountName, amountBTC := range accounts {
		fmt.Println(fmt.Sprintf(
			"Account name: %s - BTC Amount: %s", accountName, amountBTC))
	}
}

func getInfo(bitcoinClient *bitcoin.Client) {
	info, err := bitcoinClient.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf(`
Version         %d
ProtocolVersion %d
WalletVersion   %d
Balance         %f
Blocks          %d
TimeOffset      %d
Connections     %d
Proxy           %s
Difficulty      %f
TestNet         %t
KeypoolOldest   %d
KeypoolSize     %d
UnlockedUntil   %d
PaytxFee        %f
RelayFee        %f
Errors          %s
		`, info.Version, info.ProtocolVersion, info.WalletVersion, info.Balance,
		info.Blocks, info.TimeOffset, info.Connections, info.Proxy,
		info.Difficulty, info.TestNet, info.KeypoolOldest, info.KeypoolSize,
		info.UnlockedUntil, info.PaytxFee, info.RelayFee, info.Errors))
}
