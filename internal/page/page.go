package page

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/glynnforrest/kb/internal/xid"
	"gopkg.in/yaml.v3"
)

type PageType string

const (
	TYPE_NOTE   PageType = "note"
	TYPE_TASK   PageType = "task"
	TYPE_PERSON PageType = "person"
	TYPE_EVENT  PageType = "event"
)

type Page struct {
	Title    string
	ID       xid.ID
	Type     PageType
	Tags     []string
	Markdown string
}

type pageMeta struct {
	Type PageType `yaml:"type"`
	Tags []string `yaml:"tags"`
}

func (p *Page) ToString() string {
	var content strings.Builder
	content.WriteString("---\n")
	content.WriteString(fmt.Sprintf("type: %s\n", p.Type))
	content.WriteString("tags:\n")
	for _, tag := range p.Tags {
		content.WriteString(fmt.Sprintf("  - %s\n", tag))
	}
	content.WriteString("---\n\n")
	content.WriteString(fmt.Sprintf("# %s\n\n", p.Title))
	content.WriteString(p.Markdown)

	return content.String()
}

func FromFile(path string) (*Page, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	page, err := fromString(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", path, err)
	}

	id, err := xid.FromPath(path)
	if err != nil {
		return nil, fmt.Errorf("error parsing page ID: %w", err)
	}
	page.ID = id

	return page, nil
}

func fromString(content string) (*Page, error) {
	if !strings.HasPrefix(content, "---\n") {
		return nil, errors.New("missing opening frontmatter delimiter")
	}

	parts := strings.SplitN(content, "\n---\n", 2)
	if len(parts) < 2 {
		return nil, errors.New("missing closing frontmatter delimiter")
	}

	frontmatterRaw := strings.TrimPrefix(parts[0], "---\n")
	markdownBody := parts[1]

	var meta pageMeta
	if err := yaml.Unmarshal([]byte(frontmatterRaw), &meta); err != nil {
		return nil, fmt.Errorf("failed to parse YAML frontmatter: %w", err)
	}

	var title string
	lines := strings.Split(markdownBody, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			title = strings.TrimPrefix(trimmed, "# ")
			break
		}
	}

	return &Page{
		Title:    title,
		Type:     meta.Type,
		Tags:     meta.Tags,
		Markdown: strings.TrimSpace(markdownBody),
	}, nil
}
