package main

import "fmt"
import "log"

import "github.com/google/go-github/github"

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
    processPr(&pr)
  }
}

func processPr(pr *github.PullRequest) {
  fmt.Println(*pr.Title)
}
