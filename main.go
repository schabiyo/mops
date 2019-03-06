package main

import (
	"os"

	"github.com/schabiyo/mops/cmd"
	"gopkg.in/urfave/cli.v2"
)

var version string

func main() {

	app := cli.App{
		Name:        "mops",
		Usage:       "A CLI for MongoDB Ops Manager",
		Description: "MongoDB Ops Manager client ",
		Version:     version,
		Commands:    []*cli.Command{cmd.SnapCmd, cmd.ProjCmd, cmd.RestoreCmd},
	}
	app.Run(os.Args)
}
