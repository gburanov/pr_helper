package main

import (
  "log"
  "strings"
  "net/http"
  "github.com/google/go-github/github"
  "github.com/fatih/color"
)

type Repository struct {
  Organization string;
  Project string;
  Client *github.Client;
}

func NewRepository(organization string, project string, auth_token *http.Client) *Repository {
  repo := new(Repository)
  repo.Organization = organization
  repo.Project = project
  repo.Client = github.NewClient(auth_token)
  return repo
}

func (repo *Repository) listPRs() {
  prs, _, err := repo.Client.PullRequests.List(repo.Organization, repo.Project, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, pr := range prs {
    displayPR(repo.Client, *pr.Number)
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
