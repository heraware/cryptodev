// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	bitcoin "github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/heraware/cryptodev/clients"
	"github.com/spf13/cobra"
)

// litecoinCmd represents the litecoin command
var litecoinCmd = &cobra.Command{
	Use:   "litecoin",
	Short: "Run RPC command over Litecoin Node (getinfo, getnewaddress, listaccounts)",
	Run: func(cmd *cobra.Command, args []string) {
		litecoinClient := clients.NewLitecoinClient()

		if len(args) == 0 {
			log.Fatalln(`Action not provided.
Actions list:
 - getinfo
 - getnewaddress <ACCOUNT>
 - listaccounts
 - newblocks <AMOUNT OF NEW BLOCKS>
 - send <ADDRESS> <AMOUNT OF LTC>`)
		}
		switch args[0] {
		case "getinfo":
			getInfoLTC(litecoinClient)
		case "getnewaddress":
			getNewAddressLTC(litecoinClient, &args)
		case "listaccounts":
			listAccountsLTC(litecoinClient)
		case "newblocks":
			newBlocksLTC(litecoinClient, &args)
		case "send":
			sendLTC(litecoinClient, &args)
		default:
			log.Fatal(fmt.Sprintf("Action %s is not valid.", args[0]))
		}
	},
}

func init() {
	RootCmd.AddCommand(litecoinCmd)
}

func getNewAddressLTC(litecoinClient *bitcoin.Client, args *[]string) {
	account := ""
	if len(*args) > 1 {
		account = (*args)[1]
	}
	address, err := litecoinClient.GetNewAddress(account)
	if err != nil {
		log.Fatal(err)
	}
	result := fmt.Sprintf("Address: %s Account: %s", address, account)
	fmt.Println(result)
}

func listAccountsLTC(litecoinClient *bitcoin.Client) {
	accounts, err := litecoinClient.ListAccounts()
	if err != nil {
		log.Fatal(err)
	}
	var result string
	for accountName, amountBTC := range accounts {
		result = fmt.Sprintf(
			"Account name: %s - LTC Amount: %s", accountName, amountBTC)
		result = strings.Replace(result, "BTC", "LTC", 1)
		fmt.Println(result)
	}
}

func getInfoLTC(litecoinClient *bitcoin.Client) {
	info, err := litecoinClient.GetInfo()
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

func newBlocksLTC(litecoinClient *bitcoin.Client, args *[]string) {
	var amountBlocks uint64 = 1
	var err error
	if len(*args) > 1 {
		amountBlocks, err = strconv.ParseUint((*args)[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
	}
	hashes, err := litecoinClient.Generate(uint32(amountBlocks))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(hashes); i++ {
		fmt.Println(fmt.Sprintf("Block: %v", hashes[i]))
	}
}

func sendLTC(litecoinClient *bitcoin.Client, args *[]string) {
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
	litecoinClient.SendToAddress(address, amount)
}
