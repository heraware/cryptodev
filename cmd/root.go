package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the main CLI command.
var RootCmd = &cobra.Command{
	Use:   "cryptodev",
	Short: "cryptodev description",
	Long:  `long description test`,
}

// Execute calls to `Execute` RootCmd function.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
