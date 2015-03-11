package facet

import (
	"testing"
)

func TestCanClearBit(t *testing.T) {
	f := New()
	f.Clear(0)
}

func TestCanSetAndClearAndCountBits(t *testing.T) {
	f := New()
	f.Set(0)
	f.Clear(0)
	f.expectCount(t, 0)
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

func TestCanSetAndClearBigBitsAndCountBits(t *testing.T) {
	f := New()
	f.Set(10000000)
	f.Set(10000001)
	f.Clear(10000001)
	f.expectCount(t, 1)
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

func TestCanSetAllBigBitsAndClearBigBitAndCountBits(t *testing.T) {
	f := New()
	var set uint = 10000000
	f.setAllBits(set)
	f.Clear(set - 1)
	f.expectCount(t, set-1)
}

func (f *Facet) clearAllBits(start uint, end uint) {
	for i := start; i < end; i++ {
		f.Clear(i)
	}
}
