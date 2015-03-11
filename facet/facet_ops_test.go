package facet

import (
	"testing"
)

func TestCanCalcAnd(t *testing.T) {
	f, f1, f2 := New(), New(), New()
	f1.Set(20)
	f2.Set(20)
	f.And(f1, f2)
	f.expectCount(t, 1)
}
