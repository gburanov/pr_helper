package main

import (
  "log"
  "net/http"
  "github.com/google/go-github/github"
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
    repo.getPR(*pr.Number).display()
  }
}

func (repo *Repository) listPRsByLabel(label string) {
  query := "repo:" + repo.Organization + "/" + repo.Project + " label:" + label
  prs, _, err := repo.Client.Search.Issues(query, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, issue := range prs.Issues {
    repo.getPR(*issue.Number).display()
  }
}

func (repo *Repository) getPR(number int) *PR {
  pr := new(PR)
  pr.Repository = repo
  pr.Number = number
  return pr
}
