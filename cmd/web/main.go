package main

import (
  "pr_helper"

  "github.com/go-martini/martini"
)

func showPrs() string {
  ret := ""
  repo := pr_helper.RepositoryFromSettings()
  for _, pr := range repo.PRs() {
    ret += "<p>"
    ret += pr.Topic()
  }

  return ret
}

func main() {
	m := martini.Classic()
  m.Get("/", func() string {
    return showPrs()
  })
  m.Run()
}
