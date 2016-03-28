package main

import (
  "os"
  "strconv"
  "github.com/codegangsta/cli"
)

const organization = "wimdu"
const project = "wimdu"
const label = "codereview"

var verbose bool

func main() {
  verbose = false

  repo := NewRepository(organization, project, token())

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
        fillArguments(c)
        for _, pr := range repo.PRs() {
          display(filterTop(5, pr.authors()))
        }
      },
    },
    {
      Name:      "index",
      Aliases:     []string{"i"},
      Usage:     "Only PR titles",
      Action: func(c *cli.Context) {
        fillArguments(c)
        for _, pr := range repo.PRs() {
          pr.showInfo()
        }
      },
    },
    {
      Name:      "mine",
      Aliases:     []string{"m"},
      Usage:     "Mine PRs",
      Action: func(c *cli.Context) {
        fillArguments(c)
        for _, pr := range repo.myPRs() {
          display(filterTop(5, pr.authors()))
        }
      },
    },
    {
      Name:      "number",
      Aliases:     []string{"n"},
      Usage:     "PR by number",
      Action: func(c *cli.Context) {
        fillArguments(c)
        i, _ := strconv.Atoi(c.Args().First())
        repo.getPR(i).display()
      },
    },
  }

  app.Run(os.Args)
}

func fillArguments(c *cli.Context) {
  if c.Bool("verbose") == true {
    verbose = true
  }
}
