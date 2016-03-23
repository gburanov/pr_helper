package main

import "fmt"
import "log"

import "github.com/google/go-github/github"
//import "github.com/fatih/color"

const organization = "wimdu"
const project = "wimdu"

func main() {
  auth_token := token()
  client := github.NewClient(auth_token)
  listPRs(client)
}

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
  for _, file := range files {
    fileAuthors(*file.Filename)
  }
}
