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

package service

const (
	// Response messages
	MSG_VERSION = "SudaTP/1.0"
	MSG_UNKNOWN_CMD = "UNKNOWN_COMMAND"
	MSG_BYE = "TTFN"

	// Top level commands (outside transactions)
	CMD_QUIT = "QUIT"
	CMD_EXIT = "EXIT"
	CMD_PUB_KEY_REF = "PUB-KEY-REF"
	CMD_BEGIN = "BEGIN"

	// Transaction Commands (inside transactions)
	CMD_PEEK = "PEEK"
	CMD_COMMIT = "COMMIT"
	CMD_ROLLBACK = "ROLLBACK"

	// Transaction Properties
	PROP_KEY = "PublicKey"
	PROP_COMMAND = "Command"
	PROP_SIGNATURE = "Signature"
	PROP_SEP = ": "
)

var validProperties = make(map[string]bool, 4)

func init() {
	validProperties[PROP_KEY] = true
	validProperties[PROP_COMMAND] = true
	validProperties[PROP_SIGNATURE] = true
}

func IsValidProperty(prop string) bool {
	return validProperties[prop]
}
