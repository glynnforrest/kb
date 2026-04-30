package xid

import (
	"testing"
)

func TestNew(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(id) != 8 {
		t.Fatalf("expected length 8, got %d", len(id))
	}
}

func TestPath(t *testing.T) {
	id := ID("abcdefgh")
	expected := "ab/cdefgh"
	actual := id.Path()
	if expected != actual {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
