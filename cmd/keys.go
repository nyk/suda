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
	"os"
	"path"

	"github.com/nyk/suda/security/keys"
	"github.com/nyk/suda/security/keys/rsa"
	"github.com/spf13/cobra"
)

var (
	keypath string
	keyname string
	keysize int
	der     bool
)

// Perms is the default file mode for created files.
const perms os.FileMode = 0600

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
	Use:   "generate",
	Short: "Generate a private-public RSA keypair",
	Long: `The generate command generates an RSA keypair that is compatible with
	openssl and other RSA-based cryptographic software. The default format is the
	Public Encrypted Mail (PEM) format, but you can save the file in the binary
	DER format`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := rsa.GenerateKey(keysize)
		ExitOnError(err)

		basepath := path.Join(keypath, keyname)
		ExitOnError(rsa.StoreKey(keys.Private, key, basepath+".priv", perms, der))
		ExitOnError(rsa.StoreKey(keys.Public, key, basepath+".pub", perms, der))
	},
}

func init() {
	generateCmd.Flags().IntVarP(&keysize, "size", "s", 2048,
		"Set the bitsize of the generated keypair")
	generateCmd.Flags().StringVarP(&keypath, "path", "p", ".",
		"File path to store key files")
	generateCmd.Flags().StringVarP(&keyname, "name", "n", "suda-rsa",
		"Name of the key files")
	generateCmd.Flags().BoolVarP(&der, "der", "d", false,
		"Save key pair files in binary DER format")

	RootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(generateCmd)
}
