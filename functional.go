// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

// Map applies function f to every A element in input to create a new slice of
// type B.
func Map[A, B any](input []A, f func(A) B) []B {
	result := make([]B, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = f(input[i])
	}
	return result
}
