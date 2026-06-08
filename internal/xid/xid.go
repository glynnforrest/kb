package xid

import (
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

func (i ID) Path() string {
	return filepath.Join(string(i[0:2]), string(i[2:8])+".md")
}
