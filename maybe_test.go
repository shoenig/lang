// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: BSD-3-Clause

package lang

import "testing"

func TestMaybe(t *testing.T) {
	t.Parallel()

	t.Run("normal", func(t *testing.T) {
		result := Maybe(true, "alice", "bob")
		if result != "alice" {
			t.FailNow()
		}
	})

	t.Run("fallback", func(t *testing.T) {
		result := Maybe(false, "alice", "bob")
		if result != "bob" {
			t.FailNow()
		}
	})
}
