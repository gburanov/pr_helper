package main

import (
  "pr_helper"

  "github.com/go-martini/martini"
)

func showPrs() string {
  ret := ""
  repo := pr_helper.NewRepository(pr_helper.GetSettings().Organization, pr_helper.GetSettings().Project, pr_helper.Token())
  for _, pr := range repo.PRs() {
    ret += "<p>"
    ret += pr.ShowInfo()
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
