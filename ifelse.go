package lang

// If can be chained together with Else to create type-safe one-line if/else
// expressions in Go.
//
// e.g.
//
//	v := If(expression, value).Else(default)
func If[A any](b bool, v1 A) evaluation[A] {
	return evaluation[A]{
		b:  b,
		v1: v1,
	}
}

type evaluation[A any] struct {
	b  bool
	v1 A
}

func (e evaluation[A]) Else(v2 A) A {
	if e.b {
		return e.v1
	}
	return v2
}
