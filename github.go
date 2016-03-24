package main

import "log"
import "fmt"

import "github.com/google/go-github/github"

func listPRs(client *github.Client) {
  prs, _, err := client.PullRequests.List(organization, project, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, pr := range prs {
    processPr(&pr, client)
  }
}

func processPr(pr *github.PullRequest, client *github.Client) {
  fmt.Println(*pr.Title)
  files, _, err := client.PullRequests.ListFiles(organization, project, *pr.Number, nil)
  if err != nil {
    log.Fatal(err)
  }
  authors := []string{}

  for _, file := range files {
    authors = append(authors, fileAuthors(*file.Filename)...)
  }
  arrayToMap(authors)
}
