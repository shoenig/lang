// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

// Maybe returns value if b is true, or fallback otherwise.
func Maybe[T any](b bool, value T, fallback T) T {
	if b {
		return value
	}
	return fallback
}
