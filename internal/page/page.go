package page

import (
	"fmt"
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
