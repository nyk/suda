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

import (
	"errors"
	"fmt"
	"strings"
)

type Transaction struct {
	properties    map[string]string
}

func NewTransaction() *Transaction {
	return &Transaction{
		properties: make(map[string]string),
	}
}

func (trans *Transaction) Begin(sess *Session, resp *Response) error {
	for {
		reply, err := new(Request).Receive(sess)
		if err != nil {
			return err
		}

		switch (reply) {
		case CMD_PEEK:
			for key, val := range trans.properties {
				resp.Send(fmt.Sprintf("%s: %s", key, val))
			}
		case CMD_COMMIT:
			trans.Commit()
			return nil
		case CMD_ROLLBACK:
			return nil
		default:
			trans.ParseLine(reply)
		}
	}
}

func (trans *Transaction) ParseLine(line string) (bool, error) {
	keyVal := strings.SplitN(line, PROP_SEP, 2)

	if len(keyVal) < 2 {
		return false, errors.New("Malformed input in transaction")
	}

	found := trans.SetProperty(keyVal[0], keyVal[1])
	if found == false {
		return found, errors.New("Unrecognized property in transaction")
	}

	return true, nil
}

func (trans *Transaction) SetProperty(prop, value string) (bool) {
	if IsValidProperty(prop) != true {
		return false
	}
	trans.properties[prop] = value
	return true
}

func (trans *Transaction) Commit() {
}
