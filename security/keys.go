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

package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// KeyType is either Public or Private, for public and private keys.
type KeyType int8

const (
	// Private key
	Private KeyType = iota
	// Public key
	Public
	// Both public and private keys
	Both
)

func generateRsaPrivateKey(bitnum int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bitnum)
}

func storeRsaKey(keytype KeyType, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	var pkey []byte
	switch keytype {
	case Public:
		pkey, _ = x509.MarshalPKIXPublicKey(&key.PublicKey)
	case Private:
		pkey = x509.MarshalPKCS1PrivateKey(key)
	default:
		return errors.New("The key type was not set")
	}

	return ioutil.WriteFile(filepath, pkey, perm)
}

func storeRsaPemKeys(keytype KeyType, key *rsa.PrivateKey, filepath string, perm os.FileMode) {
	pemfile, _ := os.Create(filepath)

	if keytype == Private || keytype == Both {
		pem.Encode(pemfile, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
	}

	if keytype == Public || keytype == Both {
		pubkey, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)
		pem.Encode(pemfile, &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey,
		})
	}

	pemfile.Close()
	os.Chmod(filepath, perm)
}
