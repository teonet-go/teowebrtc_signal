// Copyright 2021-2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Webrtc signal server
package main

import (
	"flag"
	"log"

	"github.com/teonet-go/teowebrtc_signal"
)

var addr = flag.String("addr", ":8081", "http service address")
var addrSSL = flag.String("addr-ssl", ":8082", "https service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	teowebrtc_signal.New(*addr, *addrSSL)
}
