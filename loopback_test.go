// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package loopback

import (
	"io"
	"reflect"
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

func TestWrite(t *testing.T) {
	l := New()
	p := "Testing... 1.2.3. Is this thing on?\n"
	n, err := l.Write([]byte(p))
	if err != nil {
		t.Errorf("Write returned error %s", err)
	}
	if n != len(p) {
		t.Errorf("Write returned length %d expected %d",
			n, len(p))
	}
	n, err = l.Write([]byte(p))
	if err != nil {
		t.Errorf("Second write returned error %s", err)
	}
	if n != len(p) {
		t.Errorf("Second write returned length %d expected %d",
			n, len(p))
	}
}

func TestWriteAndRead(t *testing.T) {
	l := New()
	w := []byte("Testing... 1.2.3. Is this thing on?\n")
	n, err := l.Write(w)
	if err != nil {
		t.Errorf("Write returned error %s", err)
	}
	if n != len(w) {
		t.Errorf("Write returned length %d expected %d",
			n, len(w))
	}
	r := make([]byte, 128)
	n, err = l.Read(r)
	if err != nil {
		t.Errorf("Read returned error %s", err)
	}
	if n != len(w) {
		t.Errorf("Read returned length %d expected %d",
			n, len(w))
		return
	}
	if !reflect.DeepEqual(r[:n], w) {
		t.Errorf("Read returned %s expecting %s\n", r, w)
	}
	n, err = l.Read(r)
	if n != 0 {
		t.Errorf("Second read after write but got %d", n)
	}
	if err == nil {
		t.Errorf("Second read after write but no error")
	}
	if err != io.EOF {
		t.Errorf("Second read after write unexpected error %s", err)
	}
}

func TestWriteAndMultiRead(t *testing.T) {
	l := New()
	w := []byte("Testing... 1.2.3. Is this thing on?\n")
	n, err := l.Write(w)
	if err != nil {
		t.Errorf("Write returned error %s", err)
	}
	if n != len(w) {
		t.Errorf("Write returned length %d expected %d",
			n, len(w))
	}
	for i := 0; i < len(w); i++ {
		r := make([]byte, 1)

		n, err = l.Read(r)
		if err != nil {
			t.Errorf("Read returned error %s", err)
			return
		}

		if n != len(r) {
			t.Errorf("Read returned length %d expected %d",
				n, len(r))
			return
		}

		if r[0] != w[i] {
			t.Errorf("Read offset %d returned %c expecting %c\n",
				i, r[0], w[i])
			return
		}
	}

	r := make([]byte, 128)
	n, err = l.Read(r)
	if n != 0 {
		t.Errorf("Read after small reads should be EOF but got %d", n)
	}
	if err == nil {
		t.Errorf("Read after small reads should be EOF but no error")
	}
	if err != io.EOF {
		t.Errorf("Second read after write unexpected error %s", err)
	}
}

func TestMultiWriteAndRead(t *testing.T) {
	l := New()
	w := []byte("Testing... 1.2.3. Is this thing on?\n")
	for i := 0; i < len(w); i++ {
		n, err := l.Write(w[i : i+1])
		if err != nil {
			t.Errorf("Write returned error %s", err)
		}
		if n != 1 {
			t.Errorf("Write returned length %d expected 1", n)
		}
	}
	r := make([]byte, 128)
	n, err := l.Read(r)
	if err != nil {
		t.Errorf("Read returned error %s", err)
	}
	if n != len(w) {
		t.Errorf("Read returned length %d expected %d",
			n, len(w))
		return
	}
	if !reflect.DeepEqual(r[:n], w) {
		t.Errorf("Read returned %s expecting %s\n", r, w)
	}
	n, err = l.Read(r)
	if n != 0 {
		t.Errorf("Second read after write but got %d", n)
	}
	if err == nil {
		t.Errorf("Second read after write but no error")
	}
	if err != io.EOF {
		t.Errorf("Second read after write unexpected error %s", err)
	}
}

func TestMultiWriteAndMultiRead(t *testing.T) {
	l := New()
	w := []byte("Testing... 1.2.3. Is this thing on?\n")
	for i := 0; i < len(w); i++ {
		n, err := l.Write(w[i : i+1])
		if err != nil {
			t.Errorf("Write returned error %s", err)
		}
		if n != 1 {
			t.Errorf("Write returned length %d expected 1", n)
		}
	}
	for i := 0; i < len(w); i++ {
		r := make([]byte, 1)

		n, err := l.Read(r)
		if err != nil {
			t.Errorf("Read returned error %s", err)
			return
		}

		if n != len(r) {
			t.Errorf("Read returned length %d expected %d",
				n, len(r))
			return
		}

		if r[0] != w[i] {
			t.Errorf("Read offset %d returned %c expecting %c\n",
				i, r[0], w[i])
			return
		}
	}
	r := make([]byte, 128)
	n, err := l.Read(r)
	if n != 0 {
		t.Errorf("Read after multiread count %d", n)
	}
	if err == nil {
		t.Errorf("Read after multiread no error")
	}
	if err != io.EOF {
		t.Errorf("Read after multiread unexpected error %s", err)
	}
}
