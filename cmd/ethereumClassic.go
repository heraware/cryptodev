package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ethereumClassicCmd represents the ethereumClassic command
var ethereumClassicCmd = &cobra.Command{
	Use:   "ethereumClassic",
	Short: "Ethereum Classic Command",
	Run: func(cmd *cobra.Command, args []string) {
		// Mac and Windows support
		host := os.Getenv("DOCKER_HOST")
		if host == "" {
			host = "localhost"
		}
		fmt.Printf(`Enter to the console running the next command:

$ docker exec -it cryptodev-ethereum bash -c "geth attach http://localhost:8546"

or downloading Geth from https://github.com/ethereumproject/go-ethereum/releases and running:

$ geth attach http://%s:8546
`, host)
	},
}

func init() {
	RootCmd.AddCommand(ethereumClassicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethereumClassicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethereumClassicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
