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
