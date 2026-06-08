package cmd

import (
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	"github.com/glynnforrest/kb/internal/page"
	"github.com/glynnforrest/kb/internal/xid"
)

const dummyText = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam faucibus, lorem at luctus vulputate, nisi lectus mollis ante, sit amet finibus nulla dui at urna. Integer turpis turpis, semper suscipit tellus id, hendrerit convallis dolor. Fusce in massa nec augue fermentum vehicula. Sed et arcu lacinia, fermentum purus vitae, pharetra est. Donec diam est, ultricies et egestas vitae, aliquet eu eros. Donec non auctor velit. Sed sodales augue enim, a dapibus turpis sollicitudin nec.

Donec ac faucibus nibh. Vestibulum blandit libero eu eros posuere mattis. Ut sed purus at arcu ultricies elementum vel eu neque. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam ultricies eu nulla vitae mollis. Proin in lacus at ante maximus sollicitudin nec aliquam arcu. Suspendisse ut feugiat felis, et rutrum dolor. Proin efficitur, tortor vel pharetra efficitur, quam enim dapibus neque, eget mattis nisi quam et quam. Donec quis dolor vel tellus faucibus rutrum.

Aenean accumsan purus nulla, at rutrum lorem fringilla eu. Aliquam lacinia, quam hendrerit fringilla gravida, velit sem porttitor ex, id tristique nunc elit in libero. Nulla facilisi. Proin posuere, lacus id auctor euismod, nulla nisi aliquet tellus, sed eleifend lectus nisi porttitor nibh. Phasellus vitae felis at erat auctor placerat et in erat. Pellentesque vitae velit lorem. Integer in nunc in tortor interdum elementum eu ut sapien.`

func fixturesCmd() *command {
	c := &command{
		name: "fixtures",
	}

	var dir string
	var count int
	c.args = func(fs *flag.FlagSet) {
		fs.StringVar(&dir, "dir", "", "Path to the notes directory")
		fs.IntVar(&count, "count", 100, "Number of files to generate")
	}

	c.run = func() error {
		fmt.Printf("Generating %d fixtures in %s\n", count, dir)

		if strings.TrimSpace(dir) == "" {
			return errors.New("missing -dir")
		}

		types := []page.PageType{page.TYPE_NOTE, page.TYPE_TASK, page.TYPE_PERSON, page.TYPE_EVENT}
		tags := []string{"fixture-1", "fixture-2", "fixture-3", "fixture-4", "fixture-5"}
		dummyWords := strings.Split(dummyText, " ")

		for i := 0; i < count; i++ {
			id := xid.New()
			pagePath := filepath.Join(dir, id.Path())

			if err := os.MkdirAll(filepath.Dir(pagePath), 0755); err != nil {
				return fmt.Errorf("failed to create directory for %s: %w", pagePath, err)
			}

			p := &page.Page{
				Title:    fmt.Sprintf("Fixture Page %d", i+1),
				ID:       id,
				Type:     types[rand.IntN(len(types))],
				Tags:     []string{tags[rand.IntN(len(tags))], tags[rand.IntN(len(tags))]},
				Markdown: strings.Join(dummyWords[rand.IntN(len(dummyWords)):], " ") + "\n",
			}

			if err := os.WriteFile(pagePath, []byte(p.ToString()), 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", pagePath, err)
			}

			fmt.Printf("Created %s -> %s\n", id, pagePath)
		}

		return nil
	}

	return c
}
