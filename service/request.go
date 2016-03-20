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

package service

import "strings"

type Request struct {}

func (req Request) Receive(sess *Session) (string, error) {
	buf := make([]byte, 1024)
	bytesRead, err := sess.cx.Read(buf)
	reply := strings.TrimSpace(string(buf[0:bytesRead]))
	return reply, err
}
