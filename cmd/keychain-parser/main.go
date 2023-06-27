package main

import (
	"fmt"
	"github.com/hpcsc/keychain-parser/internal/gateway"
	"github.com/hpcsc/keychain-parser/internal/input"
	"github.com/hpcsc/keychain-parser/internal/item"
	"github.com/hpcsc/keychain-parser/internal/output"
	"github.com/hpcsc/keychain-parser/internal/updater"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
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
		Commands: []*cli.Command{
			{
				Name:  "update",
				Usage: "update to latest version",
				Action: func(*cli.Context) error {
					currentExecutable, err := os.Executable()
					if err != nil {
						return err
					}

					gw := gateway.NewGithubGateway()
					u := updater.New(runtime.GOARCH, currentExecutable, gw)
					msg, err := u.UpdateFrom(Version)
					if err != nil {
						return err
					}

					if msg != "" {
						fmt.Println(msg)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
