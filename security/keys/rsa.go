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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

// GenerateRsaKey is a wrapper function around rsa.GenerateKey function.
func GenerateRsaKey(keysize int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keysize)
}

// StoreRsaDer writes a RSA key in DER byte format to file.
func StoreRsaDer(keytype Type, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	pkey, error := EncodeRsaDer(keytype, key)
	if error != nil {
		return ErrConvert
	}
	return ioutil.WriteFile(filepath+".der", pkey, perm)
}

// StoreRsaPem writes RSA keys to file in PEM format
func StoreRsaPem(keytype Type, key *rsa.PrivateKey, filepath string, perm os.FileMode) error {
	pkey, error := EncodeRsaPem(keytype, key)
	if error != nil {
		return ErrConvert
	}
	return ioutil.WriteFile(filepath+".pem", pkey, perm)
}

// EncodeRsaDer encodes an RSA key into a slice of DER bytes
func EncodeRsaDer(keytype Type, key *rsa.PrivateKey) ([]byte, error) {
	var (
		pkey []byte
		err  error
	)

	switch keytype {
	case Private:
		pkey = x509.MarshalPKCS1PrivateKey(key)
	case Public:
		pkey, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrType
	}

	return pkey, nil
}

// EncodeRsaPem encodes an RSA key into a slice of PEM bytes
func EncodeRsaPem(keytype Type, key *rsa.PrivateKey) ([]byte, error) {
	var buf []byte

	switch keytype {
	case Private:
		buf = pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})

	case Public:
		pubkey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
		if err != nil {
			return nil, err
		}

		// TO-DO provide the option to encode
		buf = pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY", // -- RSA PUBLIC KEY -- not compatible with OpenSSL.
			Bytes: pubkey,
		})

	default:
		return nil, ErrType
	}

	return buf, nil
}
