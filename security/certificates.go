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
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

// GenerateCa function to generate a certificate authority certificate.
func signCaCertificate(template *x509.Certificate, publickey crypto.PublicKey,
	privatekey crypto.PrivateKey) ([]byte, error) {
	return x509.CreateCertificate(rand.Reader, template, template, publickey,
		privatekey)
}

func storeCaCertificate(cert []byte, filepath string, perm os.FileMode) error {
	return ioutil.WriteFile(filepath, cert, perm)
}

func makeCaTemplate() (*x509.Certificate, error) {
	// TO-DO: This needs to be populated from a configuration file.
	issuer := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Corbis"},
		OrganizationalUnit: []string{"BEN"},
		Locality:           []string{"Seattle"},
		Province:           []string{"Washington"},
	}

	// Construct the template.
	template := &x509.Certificate{
		IsCA: true,
		BasicConstraintsValid: true,
		Subject:               issuer,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 5, 5),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	return template, nil
}
