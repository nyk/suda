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
	"log"
	"net"
	"os"
)

type Session struct {
	cx  net.Conn
	key string
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		cx: conn,
	}
}

func (sess *Session) Serve() {
	defer sess.cx.Close()

	log.Printf("Client connection from %s has started.\n", sess.cx.RemoteAddr())

	resp := NewResponse(sess.cx);

	_, err := resp.Send(MSG_VERSION)
	if err != nil {
		log.Fatalf("Cannot communicate with client: %v", err)
		os.Exit(1)
	}

	for {
		reply, _ := new(Request).Receive(sess)

		switch reply {
		case CMD_QUIT:
			fallthrough
		case CMD_EXIT:
			resp.Send(MSG_BYE)
			log.Printf("Client connection from %s has ended.\n", sess.cx.RemoteAddr())
			return
		case CMD_BEGIN:
			NewTransaction().Begin(sess, resp)
		default:
			if reply[0:10] == CMD_PUB_KEY_REF {
				//sess.loadPublicKey(reply[11:])
			}
			resp.Send(MSG_UNKNOWN_CMD)
		}
	}
}

