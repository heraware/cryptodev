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
