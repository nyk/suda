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
)

func Start() {
	log.Println("Starting suda tcp service: on localhost:4444")
	ln, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalf("Cannot bind to address/port: %v", err)
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Printf("Could not create connection: %v", err)
		}
		go NewSession(conn).Serve()
	}
}
