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
	"testing"
)

func TestMakeCaTemplate(t *testing.T) {
	template, err := makeCaTemplate()
	if err != nil {
		t.Errorf("makeCaTemplate returned an error: %v", err)
	}

	if !template.IsCA {
		t.Error("template is not certificate authority")
	}
}

func TestSignCaCertificate(t *testing.T) {
	privatekey, err := generateRsaPrivateKey(2048)
	if err != nil {
		t.Errorf("cannot generate private key: %v", err)
	}
	t.Logf("Private key: %v", privatekey)

	publickey := privatekey.Public()
	t.Logf("pubkey: %v", publickey)

	template, err := makeCaTemplate()
	t.Logf("Template: %v", template)

	if err != nil {
		t.Errorf("cannot make CA template: %v", err)
	}

	cert, err := signCaCertificate(template, publickey, crypto.PrivateKey(privatekey))
	t.Logf("Certificate: %v", cert)

	if err != nil {
		t.Errorf("cannot sign CA certificate: %v", err)
	}

	err = storeCaCertificate(cert, "ca_cert.pem", 0777)
	if err != nil {
		t.Errorf("cannot store certificate: %v", err)
	}
}
