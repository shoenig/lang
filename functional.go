// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

// Map applies function f to every A element in input to create a new slice of
// type B.
func Map[A, B any](input []A, f func(A) B) []B {
	result := make([]B, len(input))
	for i := range len(input) {
		result[i] = f(input[i])
	}
	return result
}

// Head returns up to the first count elements in the given input slice.
//
// Note that this is yields a shallow copy.
func Head[A any](input []A, count int) []A {
	size := min(len(input), count)
	return input[0:size]
}
