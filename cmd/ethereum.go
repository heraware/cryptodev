package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ethereumCmd represents the ethereum command
var ethereumCmd = &cobra.Command{
	Use:   "ethereum",
	Short: "Ethereum Commands",
	Run: func(cmd *cobra.Command, args []string) {
		// Mac and Windows support
		host := os.Getenv("DOCKER_HOST")
		if host == "" {
			host = "localhost"
		}
		fmt.Printf(`Enter to the console running the next command:

$ docker exec -it cryptodev-ethereum bash -c "geth attach http://localhost:8545"

or downloading Geth from https://geth.ethereum.org/downloads/ and running:

$ geth attach http://%s:8545
`, host)
	},
}

func init() {
	RootCmd.AddCommand(ethereumCmd)

}
