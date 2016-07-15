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

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"github.com/nyk/suda/security/keys"
)

// GenerateKey is a wrapper function around rsa.GenerateKey function.
func GenerateKey(keysize int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keysize)
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

// StoreDer writes a RSA key in DER byte format to file.
func StoreDer(keytype keys.Type, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	var (
		pkey []byte
		err  error
	)

	switch keytype {
	case keys.Private:
		pkey = x509.MarshalPKCS1PrivateKey(key)
	case keys.Public:
		pkey, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
		if err != nil {
			return err
		}
	default:
		return keys.ErrType
	}

	return ioutil.WriteFile(filepath+".der", pkey, perm)
}

// StorePem writes RSA keys to file in PEM format
func StorePem(keytype keys.Type, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	pemfile, err := os.Create(filepath + ".pem")
	defer pemfile.Close()

	if err != nil {
		return err
	}

	switch keytype {
	case keys.Private:
		pem.Encode(pemfile, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})

	case keys.Public:
		pubkey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
		if err != nil {
			return err
		}

		// TO-DO provide the option to encode
		pem.Encode(pemfile, &pem.Block{
			Type:  "PUBLIC KEY", // -- RSA PUBLIC KEY -- not compatible with OpenSSL.
			Bytes: pubkey,
		})

	default:
		return keys.ErrType
	}

	os.Chmod(filepath, perm)
	return nil
}
