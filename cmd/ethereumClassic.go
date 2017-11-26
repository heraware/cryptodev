// Copyright © 2017 Yohan Graterol yohangraterol92@gmail.com
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
