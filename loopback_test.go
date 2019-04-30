// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package loopback

import (
	"io"
	"testing"
)

func TestEOFReadWithoutWrite(t *testing.T) {
	l := New()
	p := make([]byte, 128)
	n, err := l.Read(p)
	if n != 0 {
		t.Errorf("Read without write but got %d", n)
	}
	if err == nil {
		t.Errorf("Read without write but no error")
	}
	if err != io.EOF {
		t.Errorf("Read without write unexpected error %s", err)
	}
	n, err = l.Read(p)
	if n != 0 {
		t.Errorf("Second read without write but got %d", n)
	}
	if err == nil {
		t.Errorf("Second read without write but no error")
	}
	if err != io.EOF {
		t.Errorf("Second read without write unexpected error %s", err)
	}
}
