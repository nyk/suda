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

	"github.com/nyk/suda/security"
	"github.com/spf13/cobra"
)

var isCA = false

// certCmd represents the cert command
var certCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Manage certificates used for digital signatures.",
	Long: `The Suda service Listens on a TCP/IP port for client requests
that send digitally signed commands to the shell with elevated privileges.
This allows web services to securely execute commands on the shell without
having to run the web server with corresponding system privileges.`,
	Run: func(cmd *cobra.Command, args []string) {
		issuer := security.PkixName(os.Stdin)
		template, _ := security.MakeTemplate(issuer, publickey, isCA)
	},
}

func init() {
	certCmd.Flags().BoolVarP(&isCA, "ca", "a", true, "the certificate is a certificate authority")
	RootCmd.AddCommand(certCmd)
}
