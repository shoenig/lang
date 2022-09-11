package lang

import (
	"testing"
)

func eq[C comparable](t *testing.T, exp, value C) {
	if value != exp {
		t.Fatalf("not equal, expect %v, got %v", exp, value)
	}
}

func TestIfElse_Basic(t *testing.T) {
	s1 := If(true, "one").Else("two")
	eq(t, "one", s1)

	s2 := If(false, "one").Else("two")
	eq(t, "two", s2)

	i1 := If(true, 42).Else(666)
	eq(t, 42, i1)

	i2 := If(false, 42).Else(666)
	eq(t, 666, i2)
}
