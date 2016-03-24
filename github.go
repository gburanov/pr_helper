package main

import "log"

import "github.com/google/go-github/github"
import "github.com/fatih/color"

func listPRs(client *github.Client) {
  prs, _, err := client.PullRequests.List(organization, project, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, pr := range prs {
    processPr(client, *pr.Number)
  }
}

func showPrInfo(client *github.Client, num int) {
  pr, _, err := client.PullRequests.Get(organization, project, num)
  if err != nil {
    log.Fatal(err)
  }
  red := color.New(color.FgRed, color.Bold)
  red.Println(*pr.Title)
}

func processPr(client *github.Client, num int) {
  showPrInfo(client, num)
  files, _, err := client.PullRequests.ListFiles(organization, project, num, nil)
  if err != nil {
    log.Fatal(err)
  }
  authors := []string{}

  for _, file := range files {
    authors = append(authors, fileAuthors(*file.Filename)...)
  }
  arrayToMap(authors)
}
