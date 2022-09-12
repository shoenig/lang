package lang

import (
	"fmt"
)

// Combination2 represents 2 things that belong together.
type Combination2[A, B any] struct {
	A A
	B B
}

func (c Combination2[A, B]) String() string {
	return fmt.Sprintf("{%v %v}", c.A, c.B)
}

// Combine2 combines 2 things that belong together.
//
// Other languages might call this a Tuple.
func Combine2[A, B any](a A, b B) Combination2[A, B] {
	return Combination2[A, B]{A: a, B: b}
}

// Combination3 represents 3 things that belong together.
type Combination3[A, B, C any] struct {
	A A
	B B
	C C
}

func (c Combination3[A, B, C]) String() string {
	return fmt.Sprintf("{%v %v %v}", c.A, c.B, c.C)
}

// Combine3 combines 3 things that belong together.
func Combine3[A, B, C any](a A, b B, c C) Combination3[A, B, C] {
	return Combination3[A, B, C]{A: a, B: b, C: c}
}

// Combination4 represents 4 things that belong together.
type Combination4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

func (c Combination4[A, B, C, D]) String() string {
	return fmt.Sprintf("{%v %v %v %v}", c.A, c.B, c.C, c.D)
}

// Combine4 combines 4 things that belong together.
func Combine4[A, B, C, D any](a A, b B, c C, d D) Combination4[A, B, C, D] {
	return Combination4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// Combination5 represents 5 things that belong together.
type Combination5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func (c Combination5[A, B, C, D, E]) String() string {
	return fmt.Sprintf("{%v %v %v %v %v}", c.A, c.B, c.C, c.D, c.E)
}

// Combine5 combines 5 things that belong together.
func Combine5[A, B, C, D, E any](a A, b B, c C, d D, e E) Combination5[A, B, C, D, E] {
	return Combination5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}
