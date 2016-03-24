package main

import "os"

import "strconv"

import "github.com/google/go-github/github"
import "github.com/codegangsta/cli"

func main() {
  app := cli.NewApp()
  app.Name = "pr_helper"
  app.Usage = "Helps to find correct gut for PR!"
  app.Commands = []cli.Command{
    {
      Name:      "all",
      Aliases:     []string{"a"},
      Usage:     "All PRs",
      Action: func(c *cli.Context) {
        auth_token := token()
        client := github.NewClient(auth_token)
        listPRs(client)
      },
    },
    {
      Name:      "number",
      Aliases:     []string{"n"},
      Usage:     "PR by number",
      Action: func(c *cli.Context) {
        auth_token := token()
        client := github.NewClient(auth_token)
        i, _ := strconv.Atoi(c.Args().First())
        displayPR(client, i)
      },
    },
  }

  app.Run(os.Args)
}
