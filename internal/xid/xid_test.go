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

func TestFromPath(t *testing.T) {
	for _, tc := range []struct {
		given string
		expected string
		wantErr bool
	}{
		{
			given: "",
			wantErr: true,
		},
		{
			given: "abcd",
			wantErr: true,
		},
		{
			given: "abcdefgh.md",
			wantErr: true,
		},
		{
			given: "ab/cdefgh.md",
			expected: "abcdefgh",
		},
		{
			given: "/path/to/notes/ab/cdefgh.md",
			expected: "abcdefgh",
		},
		{
			given: "path/to/notes/ab/cdefgh.md",
			expected: "abcdefgh",
		},
		{
			given: "path/to/notes/ab/short.md",
			wantErr: true,
		},
	} {
		t.Run(tc.given, func(t *testing.T) {
			actual, err := FromPath(tc.given)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if tc.expected != string(actual) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestPath(t *testing.T) {
	for _, tc := range []struct {
		given string
		expected string
	}{
		{
			given: "abcdefgh",
			expected: "ab/cdefgh.md",
		},
		{
			given: "a1b2c3d4",
			expected: "a1/b2c3d4.md",
		},
	} {
		t.Run(tc.given, func(t *testing.T) {
			id := ID(tc.given)
			actual := id.Path()
			if tc.expected != actual {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
