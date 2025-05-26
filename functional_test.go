// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

import (
	"slices"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		input := []int{}
		exp := []string{}

		result := Map(input, strconv.Itoa)

		if !slices.Equal(exp, result) {
			t.FailNow()
		}
	})

	t.Run("full", func(t *testing.T) {
		input := []int{3, 8, 1}
		exp := []string{"3", "8", "1"}

		result := Map(input, strconv.Itoa)

		if !slices.Equal(exp, result) {
			t.FailNow()
		}
	})
}

func TestHead(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		input := []string{}
		result := Head(input, 3)
		if len(result) != 0 {
			t.FailNow()
		}
	})

	t.Run("shorter", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e"}
		result := Head(input, 3)
		if len(result) != 3 {
			t.FailNow()
		}
		if result[0] != "a" || result[1] != "b" || result[2] != "c" {
			t.FailNow()
		}
	})

	t.Run("longer", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		result := Head(input, 5)
		if len(result) != 3 {
			t.FailNow()
		}
		if result[0] != "a" || result[1] != "b" || result[2] != "c" {
			t.FailNow()
		}
	})
}
