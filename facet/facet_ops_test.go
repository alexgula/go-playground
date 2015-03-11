package facet

import (
	"testing"
)

func TestCanCalcAndOnEmpty(t *testing.T) {
	f, o := New(), New()
	f.And(o)
	f.expectCount(t, 0)
}

func TestCanCalcOrOnEmpty(t *testing.T) {
	f, o := New(), New()
	f.Or(o)
	f.expectCount(t, 0)
}

func TestSetAndCanCalcAndOnEqualFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(20)
	f.And(o)
	f.expectCount(t, 1)
}

func TestSetAndCanCalcAndOnBiggerFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(20)
	o.Set(40)
	f.And(o)
	f.expectCount(t, 1)
}

func TestSetAndCanCalcAndOnSmallerFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	f.Set(40)
	o.Set(20)
	f.And(o)
	f.expectCount(t, 1)
}

func TestSetAndCanCalcOrOnEqualFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(20)
	f.Or(o)
	f.expectCount(t, 1)
}

func TestSetAndCanCalcOrOnBiggerFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(20)
	o.Set(40)
	f.Or(o)
	f.expectCount(t, 2)
}

func TestSetAndCanCalcOrOnSmallerFacet(t *testing.T) {
	f, o := New(), New()
	f.Set(40)
	o.Set(20)
	f.Or(o)
	f.expectCount(t, 2)
}
