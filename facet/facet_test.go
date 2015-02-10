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
	var expected uint = 1
	if cnt := facet.Count(); cnt != expected {
		t.Fatalf("Expected facet to hold %v bit, got %v instead", expected, cnt)
	}
}

func TestCanSetBiggerBitsAndCountBits(t *testing.T) {
	facet := New()
	facet.Set(0)
	facet.Set(0)
	facet.Set(8)
	var expected uint = 2
	if cnt := facet.Count(); cnt != expected {
		t.Fatalf("Expected facet to hold %v bits, got %v instead", expected, cnt)
	}
}

func TestCanSetAllBiggerBitsAndCountBits(t *testing.T) {
	facet := New()
	var expected uint = 32
	facet.setAllBits(expected)
	if cnt := facet.Count(); cnt != expected {
		t.Fatalf("Expected facet to hold %v bits, got %v instead", expected, cnt)
	}
}

func (facet *Facet) setAllBits(count uint) {
	for i := uint(0); i < count; i++ {
		facet.Set(i)
	}
}
