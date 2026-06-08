package xid

import (
	"testing"
)

func TestNew(t *testing.T) {
	id := New()
	if len(id) != 8 {
		t.Fatalf("expected length 8, got %d", len(id))
	}
}

func TestPath(t *testing.T) {
	id := ID("abcdefgh")
	expected := "ab/cdefgh.md"
	actual := id.Path()
	if expected != actual {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
