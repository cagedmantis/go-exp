// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import (
	"testing"
)

func TestMax(t *testing.T) {
	if got := Max(1, 5); got != 5 {
		t.Fatalf("Max(1,5) != %d; want %d", got, 5)
	}
}

func TestMin(t *testing.T) {
	if got := Min(1, 5); got != 1 {
		t.Fatalf("Min(1,5) != %d; want %d", got, 1)
	}
}
