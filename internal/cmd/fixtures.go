package cmd

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/glynnforrest/kb/internal/xid"
)

func fixturesCmd() *command {
	c := &command{
		name: "fixtures",
	}

	var path string
	var count int
	c.args = func(fs *flag.FlagSet) {
		fs.StringVar(&path, "path", "", "Path to the directory where fixtures will be generated")
		fs.IntVar(&count, "count", 100, "Number of files to generate")
	}

	c.run = func() error {
		fmt.Printf("Generating %d fixtures in %s\n", count, path)

		if strings.TrimSpace(path) == "" {
			return errors.New("missing -path")
		}

		for range count {
			id := xid.New()
			pagePath := filepath.Join(path, id.Path())
			fmt.Printf("Creating %s in %s\n", id, pagePath)
		}
		return nil
	}

	return c
}
