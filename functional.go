// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

import "slices"

// Map applies function f to every A element in input to create a new slice of
// type B.
func Map[A, B any](input []A, f func(A) B) []B {
	result := make([]B, len(input))
	for i := range len(input) {
		result[i] = f(input[i])
	}
	return result
}

// FilterFunc uses f to remove elements from input.
// If f returns true for an element, that element is not included in the result.
//
// Note that FilterFunc creates a deep copy.
func FilterFunc[A any](input []A, f func(A) bool) []A {
	result := make([]A, 0, len(input))
	for _, item := range input {
		if !f(item) {
			result = append(result, item)
		}
	}
	return slices.Clip(result)
}

// Filter removes the list of items from input.
//
// Note that Filter yields a deep copy.
func Filter[A comparable](input []A, items ...A) []A {
	result := make([]A, 0, len(input))
	for _, item := range input {
		if !slices.Contains(items, item) {
			result = append(result, item)
		}
	}
	return slices.Clip(result)
}

// Head returns up to the first count elements in the given input slice.
//
// Note that this is yields a shallow copy.
func Head[A any](input []A, count int) []A {
	size := min(len(input), count)
	return input[0:size]
}

// Tail returns up to the last count elements in the given input slice.
//
// Note that this yields a shallow copy.
func Tail[A any](input []A, count int) []A {
	size := min(len(input), count)
	return input[len(input)-size:]
}
