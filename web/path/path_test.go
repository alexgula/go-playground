package path

import (
	"testing"
)

func TestPopPrefixEmpty(t *testing.T) {
	testPopPrefix(t, "", "", "")
}

func TestPopPrefixSimple(t *testing.T) {
	testPopPrefix(t, "text", "", "text")
}

func TestPopPrefixOneSegment(t *testing.T) {
	testPopPrefix(t, "/text", "/text", "")
}

func TestPopPrefixTwoSegments(t *testing.T) {
	testPopPrefix(t, "/text/rest", "/text", "/rest")
}

func testPopPrefix(t *testing.T, path, firstExpected, restExpected string) {
	first, rest := PopPrefix(path)

	if first != firstExpected {
		t.Errorf("First: value %v, expected %v", first, firstExpected)
	}

	if rest != restExpected {
		t.Errorf("Rest: value %v, expected %v", rest, restExpected)
	}
}
