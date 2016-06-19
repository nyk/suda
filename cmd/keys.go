// Copyright © 2016 Nicholas J. Cowham <nykcowham@gmail.com>
//
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

	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage public and private keys (RSA)",
	Long: `These commands allow you to create public and private key pairs
	and are provided as a convenience, so that you do not have to install
	additional dependencies, such as openssl, to perform this task`,
}

// keysCmd represents the keys command
var generateCmd = &cobra.Command{
	Use:   "generate <filepath>",
	Short: "Generate a private-public RSA keypair",
	Long: `These commands allow you to create public and private key pairs
	and are provided as a convenience, so that you do not have to install
	additional dependencies, such as openssl`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Printf("keys called: %v", args[0])
	},
}

func init() {
	RootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(generateCmd)
}
