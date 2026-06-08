package xid

import (
	"fmt"
	"strings"

	"math/rand/v2"
	"path/filepath"
)

const (
	alphabet = "0123456789abcdefghjkmnpqrstvwxyz"
	length   = 8
)

type ID string

func New() ID {
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = alphabet[rand.IntN(len(alphabet))]
	}
	return ID(id)
}

func FromPath(path string) (ID, error) {
	pathNoExt := strings.TrimSuffix(path, ".md")
	if len(pathNoExt) < 9 {
		return ID(""), fmt.Errorf("path is too short: %s", path)
	}
	last9 := pathNoExt[len(pathNoExt)-9:]
	last8 := strings.ReplaceAll(last9, string(filepath.Separator), "")

	if len(last8) != 8 {
		return ID(""), fmt.Errorf("parsed ID from path %s should be 8 chars, got: %s", path, last8)
	}

	return ID(last8), nil
}

func (i ID) Path() string {
	return filepath.Join(string(i[0:2]), string(i[2:8])+".md")
}
