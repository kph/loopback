// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package loopback provides io.Writer and io.Reader compatible interfaces
// that loop back output to input.
package loopback

import (
	"errors"
	"io"
)

// Loopback holds the pending loopback data
type Loopback struct {
	b []byte // Bytes buffered for I/O
}

// NotImplemented is the error returned for methods or options
// that are not implemented.
var NotImplemented = errors.New("Unimplemented")

// New creates a new loopback interface.
func New() (l *Loopback) {
	l = &Loopback{b: make([]byte, 0)}
	return
}

// Read returns data which has been previous written to the loopback.
func (l *Loopback) Read(p []byte) (n int, err error) {
	if len(l.b) == 0 {
		return 0, io.EOF
	}
	b := copy(p, l.b)
	l.b = l.b[b:]
	return b, nil
}

// Write is used to provide data to be looped back for later read
// operations.
func (l *Loopback) Write(p []byte) (n int, err error) {
	l.b = append(l.b, p...)
	return len(p), nil
}
