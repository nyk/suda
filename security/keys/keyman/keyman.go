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

package keyman

import (
	"encoding/pem"
	"io/ioutil"
	"os"

	"crypto/rand"
	"crypto/rsa"

	"github.com/nyk/suda/security/keys"
	rsaMan "github.com/nyk/suda/security/keys/rsa"
)

// GenerateRsaKey is a wrapper function around rsa.GenerateKey function.
func GenerateRsaKey(keysize int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keysize)
}

// StoreRsa is a dispatch function for storing RSA keys in DER or PEM format
func StoreRsa(keytype keys.Type, key *rsa.PrivateKey, filepath string, perm os.FileMode, der bool) error {
	if der {
		return rsaMan.StoreDer(keytype, key, filepath, perm)
	}

	return rsaMan.StorePem(keytype, key, filepath, perm)
}

// ParsePemFile reads and marshals a key from an io.Reader
func ParsePemFile(filepath string) ([]*pem.Block, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	blocks := make([]*pem.Block, 0, 5)

	for block, rest := pem.Decode([]byte(data)); block != nil; block, rest = pem.Decode(rest) {
		blocks = append(blocks, block)
	}

	return blocks, nil
}
