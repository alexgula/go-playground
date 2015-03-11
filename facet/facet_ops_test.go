package facet

import (
	"testing"
)

func TestCanCalcAndOnEmpty(t *testing.T) {
	f, o := New(), New()
	f.And(o)
	f.expectCount(t, 0)
}

func TestSetAndCanCalcAnd(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(20)
	f.And(o)
	f.expectCount(t, 1)
}
