package channel_test

import (
	"github.com/alexgula/go-playground/github.com/alexgula/go-playground/channel"
	"testing"
)

func TestRange(t *testing.T) {
	expected := 0
	for actual := range channel.Range(5) {
		if actual != expected {
			t.Errorf("Value %v, expected %v", actual, expected)
		}
		expected += 1
	}
}
