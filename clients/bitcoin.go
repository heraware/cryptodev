package clients

import (
	"fmt"

	bitcoin "github.com/btcsuite/btcd/rpcclient"
)

var bitcoinConnConfig = bitcoin.ConnConfig{
	Host:         "localhost:20001",
	User:         "cryptodev",
	Pass:         "cryptodev",
	DisableTLS:   true,
	HTTPPostMode: true,
}

// NewBitcoinClient connect and create bitcoinClient
func NewBitcoinClient(port uint) *bitcoin.Client {
	if port != 0 {
		bitcoinConnConfig.Host = fmt.Sprintf("localhost:%d", port)
	}
	client, err := bitcoin.New(&bitcoinConnConfig, nil)
	if err != nil {
		panic(fmt.Sprintf("[Bitcoin module] An error has ocurred: %s", err))
	}
	return client
}

// NewLitecoinClient - Litecoin is compatible with Bitcoin's API
func NewLitecoinClient() *bitcoin.Client {
	bitcoinConnConfig.Host = "localhost:21001"
	client, err := bitcoin.New(&bitcoinConnConfig, nil)
	if err != nil {
		panic(fmt.Sprintf("[Litecoin module] An error has ocurred: %s", err))
	}
	return client
}
