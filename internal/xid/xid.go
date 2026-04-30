package xid

import (
	"crypto/rand"
	"path/filepath"
)

const (
	alphabet = "0123456789abcdefghjkmnpqrstvwxyz"
	length   = 8
)

type ID string

func New() (ID, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = alphabet[int(bytes[i])%len(alphabet)]
	}

	return ID(id), nil
}

func (i ID) Path() string {
	return filepath.Join(string(i[0:2]), string(i[2:8]))
}
