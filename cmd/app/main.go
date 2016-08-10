package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/gburanov/pr_helper/lib"
)

func main() {
	manager := pr_helper.NewManager()
	repo := manager.GetRepository("wimdu", "wimdu")

	app := cli.NewApp()
	app.Name = "pr_helper"
	app.Usage = "Helps to find correct reviewer for PR!"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "Verbose mode",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "All PRs",
			Action: func(c *cli.Context) {
				for _, pr := range repo.PRs() {
					displayPR(pr)
					fmt.Println()
				}
			},
		},
		{
			Name:    "index",
			Aliases: []string{"i"},
			Usage:   "Only PR titles",
			Action: func(c *cli.Context) {
				for _, pr := range repo.PRs() {
					pr.Topic()
				}
			},
		},
		{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "PR by number",
			Action: func(c *cli.Context) {
				i, _ := strconv.Atoi(c.Args().First())
				pr, err := repo.GetPR(i)
				if err != nil {
					log.Fatal(err)
				}
				displayPR(*pr)
			},
		},
		{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "PR by URL",
			Action: func(c *cli.Context) {
				url := c.Args().First()
				pr, err := manager.GetPR(url)
				if err != nil {
					log.Fatal(err)
				}
				displayPR(*pr)
			},
		},
	}

	app.Run(os.Args)
}
