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
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/nyk/suda/security/keys"
)

// SignCaCertificate self-signs the certificate authority certificate.
func SignCaCertificate(template *x509.Certificate, publickey crypto.PublicKey,
	privatekey crypto.PrivateKey) ([]byte, error) {
	return x509.CreateCertificate(rand.Reader, template, template, publickey,
		privatekey)
}

// StoreCertificate stores a certificate to the file system.
func StoreCertificate(cert []byte, filepath string, perm os.FileMode) error {
	return ioutil.WriteFile(filepath, cert, perm)
}

// MakeTemplate constructs a certificate authority template with public key.
func MakeTemplate(subject *pkix.Name, publickey *crypto.PublicKey, isCA bool) (*x509.Certificate, error) {
	key, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return nil, keys.ErrConvert
	}

	// SKI is supposed to be an SHA1 hash of the public key of the subject.
	subjectKeyID := sha1.Sum(key)

	// Construct the template.
	template := &x509.Certificate{
		IsCA: isCA,
		BasicConstraintsValid: true,
		Subject:               *subject,
		SubjectKeyId:          subjectKeyID[:],
		SerialNumber:          big.NewInt(1234),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 5, 5),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	return template, nil
}

// PkixName creates a DN record populated from an io.Reader
func PkixName(stream io.Reader) (name *pkix.Name) {
	r := bufio.NewReader(stream)
	name = &pkix.Name{
		Country:            PkixField(r, "Country"),
		Organization:       PkixField(r, "Organization"),
		OrganizationalUnit: PkixField(r, "Organizational Unit"),
		Locality:           PkixField(r, "Locality/city"),
		Province:           PkixField(r, "Province"),
		StreetAddress:      PkixField(r, "Street Address"),
		PostalCode:         PkixField(r, "Postal Code"),
		CommonName:         PkixField(r, "Common Name")[0],
	}
	return
}

// PkixField constructs a DN field from a bufio reader.
func PkixField(r *bufio.Reader, label string) []string {
	fmt.Printf("%s: ", label)
	field, _ := r.ReadString('\n')
	return []string{strings.TrimSpace(field)}
}
