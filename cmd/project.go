package cmd

import (
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"github.com/schabiyo/mops/apiclient"
	"gopkg.in/urfave/cli.v2"
)

var (
	ProjCmd = &cli.Command{
		Name:    "project",
		Aliases: []string{"s"},
		Usage:   "options for project management",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all projects the user has access to",
				Action: func(c *cli.Context) error {

					client := apiclient.NewOpsManagerAPI()
					var page int
					page = 1
					printer := tableprinter.New(os.Stdout)
					for {
						projs, _ := client.ProjectAPI.GetProjects(client, page)
						//For each project get the clusters
						if len(projs.Projects) == 0 {
							fmt.Println(projs.TotalCount, " project(s) found")
							break
						}
						page++
						client.ProjectAPI.PrintResultTable(printer, projs.Projects)
						fmt.Println(">> For additional result press any key or 'q' to stop")
						var s string
						fmt.Scan(&s)
						if s == "q" {
							break
						}
					}
					return nil
				},
			},
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
