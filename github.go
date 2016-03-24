package main

import "log"
import "strings"

import "github.com/google/go-github/github"
import "github.com/fatih/color"

func listPRs(client *github.Client) {
  prs, _, err := client.PullRequests.List(organization, project, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, pr := range prs {
    authors := processPr(client, *pr.Number)
    authors = filterTop(5, authors)
    display(authors)
  }
}

func showPrInfo(client *github.Client, num int) {
  pr, _, err := client.PullRequests.Get(organization, project, num)
  if err != nil {
    log.Fatal(err)
  }
  red := color.New(color.FgRed, color.Bold)
  red.Println(*pr.Title, "#", num)
}

func processPr(client *github.Client, num int) map[string]int {
  showPrInfo(client, num)
  files, _, err := client.PullRequests.ListFiles(organization, project, num, nil)
  if err != nil {
    log.Fatal(err)
  }

  results := 0
  authors_channel := make(chan []string)
  for _, file := range files {
    if strings.HasPrefix(*file.Filename, "phrase") {
      continue
    }
    results++
    go func(fileName string) {
      authors_channel <- fileAuthors(fileName)
    }(*file.Filename)
  }
  authors := []string{}
  for i := 1; i <= results; i++ {
    authors = append(authors, <-authors_channel...)
  }
  return arrayToMap(authors)
}
