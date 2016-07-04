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

// ErrKeyType represents an invalid or empty key type value.
var ErrKeyType = errors.New("Invalid key type")

// GenerateKey is a wrapper function around rsa.GenerateKey function.
func GenerateKey(keysize uint16) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keysize)
}

// StoreKey writes a RSA key in DER byte format to file.
func StoreKey(keytype KeyType, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	var pkey []byte
	switch keytype {
	case Public:
		pkey, _ = x509.MarshalPKIXPublicKey(&key.PublicKey)
	case Private:
		pkey = x509.MarshalPKCS1PrivateKey(key)
	default:
		return ErrKeyType
	}

	return ioutil.WriteFile(filepath, pkey, perm)
}

// StoreKeysPem writes RSA keys to file in PEM format
func StoreKeysPem(keytype KeyType, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	pemfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

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
	return nil
}
