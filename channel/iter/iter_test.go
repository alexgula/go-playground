package iter

import (
	"testing"
)

func TestRange(t *testing.T) {
	expected := 0
	for actual := range Range(5) {
		if actual != expected {
			t.Errorf("Value %v, expected %v", actual, expected)
		}
		expected += 1
	}
}
