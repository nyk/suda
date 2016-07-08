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
	"crypto/rsa"
	"crypto/x509/pkix"
	"testing"
)

var privatekey, keyErr = rsa.GenerateKey(rand.Reader, 2048)

var name = &pkix.Name{
	Country:            []string{"UK"},
	Organization:       []string{"Suda Project"},
	OrganizationalUnit: []string{"Dev"},
}

func init() {
	if keyErr != nil {
		panic("Cannot generate private key")
	}
}

func TestMakeCaTemplate(t *testing.T) {
	publickey := privatekey.Public()
	template, err := MakeTemplate(name, &publickey, false)
	if err != nil {
		t.Errorf("makeCaTemplate returned an error: %v", err)
	}

	if !template.IsCA {
		t.Error("template is not certificate authority")
	}
}

func TestSignCaCertificate(t *testing.T) {
	publickey := privatekey.Public()
	t.Logf("pubkey: %v", publickey)

	template, err := MakeTemplate(name, &publickey, true)
	t.Logf("Template: %v", template)

	if err != nil {
		t.Errorf("cannot make CA template: %v", err)
	}

	cert, err := SignCaCertificate(template, publickey, crypto.PrivateKey(privatekey))
	t.Logf("Certificate: %v", cert)

	if err != nil {
		t.Errorf("cannot sign CA certificate: %v", err)
	}

	err = StoreCertificate(cert, "ca_cert.crt", 0777)
	if err != nil {
		t.Errorf("cannot store certificate: %v", err)
	}
}
