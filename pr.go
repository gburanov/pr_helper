package main

import (
  "log"
  "github.com/fatih/color"
)

type PR struct {
  Repository *Repository
  Number int
}

func (pr *PR) display() {
  authors := processPr(pr.Repository.Client, pr.Number)
  authors = filterTop(5, authors)
  display(authors)
}

func (pr *PR) showInfo() {
  pr_, _, err := pr.Repository.Client.PullRequests.
    Get(pr.Repository.Organization, pr.Repository.Project, pr.Number)
  if err != nil {
    log.Fatal(err)
  }
  red := color.New(color.FgRed, color.Bold)
  red.Println(*pr_.Title, "#", pr.Number)
}
