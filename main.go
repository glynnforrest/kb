package main

import (
	"fmt"

	"github.com/glynnforrest/kb/internal/xid"
)

func main() {
	id := xid.New()
	fmt.Println(id)
	fmt.Println(id.Path())
}
