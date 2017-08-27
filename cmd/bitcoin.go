package cmd

import (
	"fmt"
	"log"
	"strconv"

	bitcoin "github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
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

func newBlocks(bitcoinClient *bitcoin.Client, args *[]string) {
	var amountBlocks uint64 = 1
	var err error
	if len(*args) > 1 {
		amountBlocks, err = strconv.ParseUint((*args)[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
	}
	hashes, err := bitcoinClient.Generate(uint32(amountBlocks))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(hashes); i++ {
		fmt.Println(fmt.Sprintf("Block: %v", hashes[i]))
	}
}

func send(bitcoinClient *bitcoin.Client, args *[]string) {
	var address btcutil.Address
	var amountFloat float64
	var amount btcutil.Amount
	var err error
	if len(*args) > 2 {
		address, err = btcutil.DecodeAddress((*args)[1], nil)
		if err != nil {
			log.Fatal(err)
		}
		amountFloat, err = strconv.ParseFloat((*args)[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		amount, err = btcutil.NewAmount(amountFloat)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Put all arguments to this action. send <BTC ADDRESS> <BTC AMOUNT>")
	}
	bitcoinClient.SendToAddress(address, amount)
}
