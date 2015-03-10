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
	f.Set(8)
	f.expectCount(t, 2)
}

func TestCanSetAndClearBiggerBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(8)
	f.Clear(8)
	f.expectCount(t, 0)
}

func TestCanSetAndDoubleClearBiggerBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(8)
	f.Clear(8)
	f.Clear(8)
	f.expectCount(t, 0)
}

func TestCanClearHugeBitsAndCountBits(t *testing.T) {
	f := New()
	f.Clear(1000000000000000)
	f.expectCount(t, 0)
}

func TestCanSetBigBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(10000000)
	f.expectCount(t, 1)
}

func TestCanSetAndClearBigBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(10000000)
	f.Set(10000001)
	f.Clear(10000001)
	f.expectCount(t, 1)
}

func TestCanSetAllBiggerBitsAndCountBits(t *testing.T) {
	f := New()
	var set uint = 32
	f.setAllBits(set)
	f.expectCount(t, set)
}

func TestCanSetAllBiggerBitsAndClearSomeBitsAndCountBits(t *testing.T) {
	f := New()
	var set uint = 32
	var clearStart uint = 8
	var clearEnd uint = 16
	f.setAllBits(set)
	f.clearAllBits(clearStart, clearEnd)
	f.expectCount(t, set-(clearEnd-clearStart))
}

func TestCanSetAllBigBitsAndCountBits(t *testing.T) {
	f := New()
	var set uint = 10000000
	f.setAllBits(set)
	f.expectCount(t, set)
}

func TestCanSetAllBigBitsAndclearbigBitAndCountBits(t *testing.T) {
	f := New()
	var set uint = 10000000
	f.setAllBits(set)
	f.Clear(set - 1)
	f.expectCount(t, set-1)
}

func (f *Facet) setAllBits(count uint) {
	for i := uint(0); i < count; i++ {
		f.Set(i)
	}
}

func (f *Facet) clearAllBits(start uint, end uint) {
	for i := start; i < end; i++ {
		f.Clear(i)
	}
}

func (f *Facet) expectCount(t *testing.T, expected uint) {
	if cnt := f.Count(); cnt != expected {
		t.Fatalf("Expected f to hold %v bits, got %v instead", expected, cnt)
	}
}
