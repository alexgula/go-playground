package facet

import (
	"testing"
)

func TestCanBeCreated(t *testing.T) {
	New()
}

func TestCanCountBits(t *testing.T) {
	f := New()
	f.Count()
}

func (f *Facet) expectCount(t *testing.T, expected uint) {
	if cnt := f.Count(); cnt != expected {
		t.Fatalf("Facet expected to hold %v bits, got %v instead", expected, cnt)
	}
}
