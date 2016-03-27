package main

import (
  "os"
  "strconv"
  "github.com/codegangsta/cli"
)

const organization = "wimdu"
const project = "wimdu"

func main() {
  repo := NewRepository(organization, project, token())

  app := cli.NewApp()
  app.Name = "pr_helper"
  app.Usage = "Helps to find correct gut for PR!"
  app.Commands = []cli.Command{
    {
      Name:      "all",
      Aliases:     []string{"a"},
      Usage:     "All PRs",
      Action: func(c *cli.Context) {
        repo.listPRs()
      },
    },
    {
      Name:      "mine",
      Aliases:     []string{"m"},
      Usage:     "Mine PRs",
      Action: func(c *cli.Context) {
        repo.listPRs()
      },
    },
    {
      Name:      "number",
      Aliases:     []string{"n"},
      Usage:     "PR by number",
      Action: func(c *cli.Context) {
        i, _ := strconv.Atoi(c.Args().First())
        repo.getPR(i).display()
      },
    },
  }

  app.Run(os.Args)
}
