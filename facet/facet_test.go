package facet

import (
	"testing"
)

func TestCanBeCreated(t *testing.T) {
	New()
}

func TestCanSetBit(t *testing.T) {
	facet := New()
	facet.Set(0)
}

func TestCanCountBits(t *testing.T) {
	facet := New()
	facet.Count()
}

func TestCanSetAndCountBits(t *testing.T) {
	facet := New()
	facet.Set(0)
	if cnt := facet.Count(); cnt != 1 {
		t.Fatalf("Expected facet to hold 1 bit, got %v instead", cnt)
	}
}

func TestCanSetSameBitAndCountBits(t *testing.T) {
	facet := New()
	facet.Set(0)
	facet.Set(0)
	if cnt := facet.Count(); cnt != 1 {
		t.Fatalf("Expected facet to hold 1 bit, got %v instead", cnt)
	}
}
