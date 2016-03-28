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

func (repo *Repository) listPRsbyQuery(query string) []PR  {
  prs, _, err := repo.Client.Search.Issues(query, nil)
  if err != nil {
    log.Fatal(err)
  }
  ret := []PR {}
  for _, pr := range prs.Issues {
    ret = append(ret, *repo.getPR(*pr.Number))
  }
  return ret
}

func (repo *Repository) myPRs() []PR {
  query := "repo:" + repo.Organization + "/" + repo.Project + " label:" + label + " author:gburanov"
  return repo.listPRsbyQuery(query)
}

func (repo *Repository) PRs() []PR {
  query := "repo:" + repo.Organization + "/" + repo.Project + " label:" + label
  return repo.listPRsbyQuery(query)
}

func (repo *Repository) getPR(number int) *PR {
  pr := new(PR)
  pr.Repository = repo
  pr.Number = number
  return pr
}
