package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const version = "2017.8"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show cryptodev CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		message := fmt.Sprintf("Cryptodev version: %s\n", version)
		fmt.Println(message)
		c := color.New(color.FgGreen, color.Bold)
		c.Println("If you want to support this cause you can donate to:")
		fmt.Println(`
ETC - 0x69fEd41D84902Ac638cd31Ca8AC5249E4E8fa241
BTC - 115GHLKuLfPh1qVjyRQFSyiYQLGJLSMdTj
LTC - LNdJQ3hM795ppTijD94PM9D1cHerK6d1AE
ETH - 0x1B5d7ceb8d9B5B56ccC963D8C2FE4EA254ABDBD2`)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
