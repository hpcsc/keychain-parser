package main

import (
	"github.com/hpcsc/keychain-parser/internal/input"
	"github.com/hpcsc/keychain-parser/internal/item"
	"github.com/hpcsc/keychain-parser/internal/output"
	"os"
)

func main() {
	lines := input.From(os.Stdin)
	items := item.From(lines)
	if err := output.AsJson(items, os.Stdout); err != nil {
		panic(err)
	}
}
