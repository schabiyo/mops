package cmd

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

var (
	OrgCmd = &cli.Command{
		Name:    "organization",
		Aliases: []string{"org"},
		Usage:   "options for organization management",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get one snapshop details",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "delete a snapshot",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
		},
	}
)
