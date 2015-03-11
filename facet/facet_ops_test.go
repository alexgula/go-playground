package facet

import (
	"testing"
)

func TestCanCalcAnd(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(20)
	f.And(o)
	f.expectCount(t, 1)
}
