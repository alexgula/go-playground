package facet

import (
	"testing"
)

func TestCanSetBit(t *testing.T) {
	f := New()
	f.Set(0)
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

func TestCanSetBiggerBitsAndCountBits(t *testing.T) {
	f := New()
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
	var set uint = 32
	f.setAllBits(set)
	f.expectCount(t, set)
}

func TestCanSetAllBigBitsAndCountBits(t *testing.T) {
	f := New()
	var set uint = 10000000
	f.setAllBits(set)
	f.expectCount(t, set)
}

func (f *Facet) setAllBits(count uint) {
	for i := uint(0); i < count; i++ {
		f.Set(i)
	}
}
