// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/tarm/serial"
)

func init() {
	configureSerialLogOutput = configureSerialLogOutputWindows
}

func configureSerialLogOutputWindows() {
	c := &serial.Config{Name: "COM1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		// Oh well, we tried. This empirically works
		// on Windows on GCE.
		// We can log here anyway and hope somebody sees it
		// in a GUI console:
		log.Printf("serial.OpenPort: %v", err)
		return
	}
	log.SetOutput(s)
}