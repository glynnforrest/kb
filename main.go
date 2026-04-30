package main

import (
	"fmt"

	"github.com/glynnforrest/kb/internal/xid"
)

func main() {
	id, err := xid.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	fmt.Println(id.Path())
}
