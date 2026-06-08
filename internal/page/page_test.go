package page

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	for _, tc := range []struct {
		name     string
		contents string
		expected *Page
		wantErr  bool
	}{
		{
			name:     "empty",
			contents: ``,
			wantErr:  true,
		},
		{
			name:     "missing frontmatter",
			contents: `# Some page`,
			wantErr:  true,
		},
		{
			name: "valid",
			contents: `---
type: note
tags:
  - fixture-1
---

# Some page

Some text here
`,
			expected: &Page{
				Type:  TYPE_NOTE,
				Tags:  []string{"fixture-1"},
				Title: "Some page",
				Markdown: `# Some page

Some text here`,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p, err := fromString(tc.contents)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !reflect.DeepEqual(tc.expected, p) {
				t.Fatalf("expected %v, got %v", tc.expected, p)
			}
		})
	}
}

func TestFromFile(t *testing.T) {
	for _, tc := range []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "missing file",
			path:    "not/a/path",
			wantErr: true,
		},
		{
			name: "valid file",
			path: "testdata/aa/123456.md",
		},
		{
			name:    "bad file contents",
			path:    "testdata/aa/abcdef.md",
			wantErr: true,
		},
		{
			name:    "bad file path",
			path:    "testdata/aa/12345.md",
			wantErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := FromFile(tc.path)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
