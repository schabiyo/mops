package cmd

import (
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"github.com/schabiyo/mops/apiclient"
	"gopkg.in/urfave/cli.v2"
)

var (
	SnapCmd = &cli.Command{
		Name:    "snapshot",
		Aliases: []string{"s"},
		Usage:   "options for snapshot management",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "list all snapshot of a given cluster",
				Action: func(c *cli.Context) error {
					if c.Args().Len() != 2 {
						fmt.Println("Project ID and Cluster ID required ")
						os.Exit(1)
					}
					fmt.Println(fmt.Printf("listing snapshots for group:%s and cluster:%s", c.Args().First(), c.Args().Get(1)))
					//"5c6d79b1432081b6c7f9abc2"
					//5c6d7ce712583ef50c817f25
					client := apiclient.NewOpsManagerAPI()
					var page int
					page = 1
					printer := tableprinter.New(os.Stdout)
					for {
						snaps, _ := client.SnapshotAPI.GetSnapshots(client, c.Args().First(), c.Args().Get(1), page)
						//For each project get the clusters
						if len(snaps.Snapshots) == 0 {
							fmt.Println(snaps.TotalCount, " snapshot(s) found")
							break
						}
						page++
						client.SnapshotAPI.PrintResultTable(printer, snaps.Snapshots)
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
