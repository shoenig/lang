package lang

import (
	"sync"
	"testing"
)

func TestCritical(t *testing.T) {
	t.Parallel()

	m := new(sync.Mutex)
	var reached bool

	Critical(m, func() {
		reached = true
	})

	if !reached {
		t.FailNow()
	}
}

func TestCriticalGet(t *testing.T) {
	t.Parallel()

	m := new(sync.Mutex)

	result := CriticalGet(m, func() int {
		return 42
	})

	if result != 42 {
		t.FailNow()
	}
}
