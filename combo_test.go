package lang

import (
	"testing"
	"time"
)

func TestCombination_2(t *testing.T) {
	c := Combine2[int, string](42, "foo")
	s := c.String()
	eq(t, "{42 foo}", s)
}

func TestCombination_3(t *testing.T) {
	c := Combine3[float32, string, int](1.1, "hi", 8)
	s := c.String()
	eq(t, "{1.1 hi 8}", s)
}

func TestCombination_4(t *testing.T) {
	c := Combine4[int, string, int, string](1, "two", 3, "four")
	s := c.String()
	eq(t, "{1 two 3 four}", s)
}

func TestCombination_5(t *testing.T) {
	now := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	c := Combine5[time.Time, int, []byte, string, func()](
		now,
		42,
		[]byte{1, 2, 3},
		"hello",
		nil,
	)
	s := c.String()
	eq(t, "{2001-02-03 04:05:06.000000007 +0000 UTC 42 [1 2 3] hello <nil>}", s)
}
