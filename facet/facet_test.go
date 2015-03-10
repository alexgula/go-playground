package facet

import (
	"testing"
)

func TestCanBeCreated(t *testing.T) {
	New()
}

func TestCanSetBit(t *testing.T) {
	f := New()
	f.Set(0)
}

func TestCanClearBit(t *testing.T) {
	f := New()
	f.Clear(0)
}

func TestCanCountBits(t *testing.T) {
	f := New()
	f.Count()
}

func TestCanSetAndCountBits(t *testing.T) {
	f := New()
	f.Set(0)
	f.expectCount(t, 1)
}

func TestCanSetSameBitAndCountBits(t *testing.T) {
	f := New()
	f.Set(0)
	f.Set(0)
	f.expectCount(t, 1)
}

func TestCanSetAndClearAndCountBits(t *testing.T) {
	f := New()
	f.Set(0)
	f.Clear(0)
	f.expectCount(t, 0)
}

func TestCanSetBiggerBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(0)
	f.Set(0)
	f.Set(8)
	f.expectCount(t, 2)
}

func TestCanSetBigBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(10000000)
	f.expectCount(t, 1)
}

func TestCanSetAllBiggerBitsAndCountBits(t *testing.T) {
	f := New()
	var expected uint = 32
	f.setAllBits(expected)
	f.expectCount(t, expected)
}

func TestCanSetAllBigBitsAndCountBits(t *testing.T) {
	f := New()
	var expected uint = 10000000
	f.setAllBits(expected)
	f.expectCount(t, expected)
}

func (f *Facet) setAllBits(count uint) {
	for i := uint(0); i < count; i++ {
		f.Set(i)
	}
}

func (f *Facet) expectCount(t *testing.T, expected uint) {
	if cnt := f.Count(); cnt != expected {
		t.Fatalf("Expected f to hold %v bits, got %v instead", expected, cnt)
	}
}
