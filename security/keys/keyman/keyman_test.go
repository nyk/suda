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

import "testing"

func TestGeneratePrivateKey(t *testing.T) {

}

func TestParsePrivatePemFile(t *testing.T) {
	blocks, _ := ParsePemFile("./test/private.pem")
	if len(blocks) != 1 {
		t.Fail()
	}
	if blocks[0].Type != "RSA PRIVATE KEY" {
		t.Fail()
		t.Logf("Block type is %s", blocks[0].Type)
	}
	t.Logf("The private.pem file contains %d blocks", len(blocks))
}

func TestParsePublicPemFile(t *testing.T) {
	blocks, _ := ParsePemFile("./test/public.pem")
	if len(blocks) != 1 {
		t.Fail()
	}
	if blocks[0].Type != "PUBLIC KEY" {
		t.Fail()
		t.Logf("Block type is %s", blocks[0].Type)
	}
	t.Logf("The public.pem file contains %d blocks", len(blocks))
}

func TestParseBothPemFile(t *testing.T) {
	blocks, _ := ParsePemFile("./test/both.pem")
	if len(blocks) != 2 {
		t.Fail()
	}
	if blocks[0].Type != "RSA PRIVATE KEY" {
		t.Fail()
		t.Logf("Block type is %s", blocks[0].Type)
	}
	if blocks[1].Type != "PUBLIC KEY" {
		t.Fail()
		t.Logf("Block type is %s", blocks[1].Type)
	}
	t.Logf("The both.pem file contains %d blocks", len(blocks))
}
