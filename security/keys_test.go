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

import "testing"

func TestGeneratePrivateKey(t *testing.T) {
	_, err := generateRsaPrivateKey(2048)
	if err != nil {
		t.Error("There was an error generating the RSA key")
	}

	//storeRsaKey(Private, key, "private.key", 0777)
	//storeRsaKey(Public, key, "public.key", 0777)
	//storeRsaPemKeys(Public, key, "public.pem", 0777)
}
