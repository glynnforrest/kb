package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/glynnforrest/kb/internal/page"
)

func lsCmd() *command {
	c := &command{
		name: "ls",
	}

	var dir string
	c.args = func(fs *flag.FlagSet) {
		fs.StringVar(&dir, "dir", "", "Path to the notes directory")
	}

	c.run = func() error {
		if strings.TrimSpace(dir) == "" {
			return errors.New("missing -dir")
		}

		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}

			p, err := page.FromFile(path)
			if err != nil {
				return err
			}

			tagsStr := fmt.Sprintf("[%s]", strings.Join(p.Tags, ", "))
			fmt.Printf("%s | %-20s | %-7s | %s\n",
				p.ID, p.Title, p.Type, tagsStr)

			return nil
		})

		if err != nil {
			return fmt.Errorf("error walking directory: %w", err)
		}

		return nil
	}

	return c
}
