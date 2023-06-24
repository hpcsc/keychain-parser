package main

import (
	"github.com/hpcsc/keychain-parser/internal/input"
	"github.com/hpcsc/keychain-parser/internal/item"
	"github.com/hpcsc/keychain-parser/internal/output"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	Version = "main"
)

func main() {
	app := &cli.App{
		Name:    "keychain-parser",
		Usage:   "parse output from MacOS security command and output in a more readable format (.e.g. JSON)",
		Version: Version,
		Action: func(*cli.Context) error {
			lines := input.From(os.Stdin)
			items := item.From(lines)
			return output.AsJson(items, os.Stdout)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
