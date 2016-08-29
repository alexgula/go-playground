package merge

import (
	"testing"
)

func TestMerge(t *testing.T) {
	out := make(chan int)
	a := make(chan int)
	b := make(chan int)

	aVal, bVal, sendCount := 1, 2, 10

	go send(a, aVal, sendCount)
	go send(b, bVal, sendCount)
	go merge(out, a, b)

	count := 0
	for v := range out {
		if v != aVal && v != bVal {
			t.Errorf("Value %v, expected %v or %v", v, aVal, bVal)
		}
		count++
	}
	if count != sendCount+sendCount {
		t.Errorf("Gor %v values, expected %v", count, sendCount+sendCount)
	}
}

func send(ch chan<- int, val, size int) {
	for i := 0; i < size; i++ {
		ch <- val
	}
	close(ch)
}
