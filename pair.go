// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

// A Pair contains two elements of any type(s).
type Pair[T, U any] struct {
	A T
	B U
}
