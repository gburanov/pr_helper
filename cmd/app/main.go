package main

import (
  "os"
  "fmt"
  "strconv"
  "github.com/codegangsta/cli"

  "pr_helper"
)

const organization = "wimdu"
const project = "wimdu"
const label = "codereview"

func main() {
  repo := pr_helper.NewRepository(organization, project, pr_helper.Token())

  app := cli.NewApp()
  app.Name = "pr_helper"
  app.Usage = "Helps to find correct gut for PR!"
  app.EnableBashCompletion = true

  app.Flags = []cli.Flag {
    cli.BoolTFlag{
      Name: "verbose",
      Usage: "Verbose mode",
    },
  }

  app.Commands = []cli.Command{
    {
      Name:      "all",
      Aliases:     []string{"a"},
      Usage:     "All PRs",
      Action: func(c *cli.Context) {
        for _, pr := range repo.PRs() {
          displayPR(&pr)
          fmt.Println()
        }
      },
    },
    {
      Name:      "index",
      Aliases:     []string{"i"},
      Usage:     "Only PR titles",
      Action: func(c *cli.Context) {
        for _, pr := range repo.PRs() {
          pr.ShowInfo()
        }
      },
    },
    {
      Name:      "mine",
      Aliases:     []string{"m"},
      Usage:     "Mine PRs",
      Action: func(c *cli.Context) {
        for _, pr := range repo.MyPRs() {
          displayPR(&pr)
        }
      },
    },
    {
      Name:      "number",
      Aliases:     []string{"n"},
      Usage:     "PR by number",
      Action: func(c *cli.Context) {
        i, _ := strconv.Atoi(c.Args().First())
        displayPR(repo.GetPR(i))
      },
    },
  }

  app.Run(os.Args)
}
