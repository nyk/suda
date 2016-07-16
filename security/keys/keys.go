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

package keys

import (
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// Type is either Public or Private, for public and private keys.
type Type int8

const (
	// Private key
	Private Type = iota
	// Public key
	Public
)

var (
	// ErrType indicates an invalid option for a key (Public/Private)
	ErrType = errors.New("Invalid key type")
	// ErrConvert indicates that the key could not be converted to bytes
	ErrConvert = errors.New("Cannot convert key to bytes")
)

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
