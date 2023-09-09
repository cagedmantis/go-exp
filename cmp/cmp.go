// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

import "golang.org/x/exp/constraints"

// Max returns the largest between a and b. If a is equal to b, a is returned.
func Max[T constraints.Ordered](a, b T) T{
	if a >= b {
		return a
	}
	return b
}

// Min returns the smallest between a and b. If a is equal to b, a is returned.
func Min[T constraints.Ordered](a, b T) T{
	if a <= b {
		return a
	}
	return b
}
