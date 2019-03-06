package cmd

import (
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"github.com/schabiyo/mops/apiclient"
	"gopkg.in/urfave/cli.v2"
)

var (
	RestoreCmd = &cli.Command{
		Name:    "restore",
		Aliases: []string{"r"},
		Usage:   "options for DB restore",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all restore jobs of a given cluster",
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 2 {
						fmt.Println("Project ID and Cluster ID required ")
						os.Exit(1)
					}
					client := apiclient.NewOpsManagerAPI()
					var page int
					page = 1
					printer := tableprinter.New(os.Stdout)
					for {
						restores, _ := client.RestoreAPI.GetRestoreJobs(client, c.Args().First(), c.Args().Get(1), page)
						//For each project get the clusters
						if len(restores.Jobs) == 0 {
							fmt.Println(restores.TotalCount, " restore job(s) found")
							break
						}
						page++
						client.RestoreAPI.PrintResultTable(printer, restores.Jobs)
					}
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create a Restore job",
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 3 {
						fmt.Println("Project ID, Cluster ID and Snapshot ID required ")
						os.Exit(1)
					}
					client := apiclient.NewOpsManagerAPI()
					res, _ := client.RestoreAPI.CreateRestoreJob(client, c.Args().First(), c.Args().Get(1), c.Args().Get(2))
					fmt.Println(res.Jobs)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "Get one Retsore Job details",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
		},
	}
)
