package main

import (
  "pr_helper"

  "log"

  "net/http"
  "html/template"
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
  m.Get("/", func(res http.ResponseWriter, req *http.Request) {
    t, err := template.ParseFiles("cmd/web/index.gtpl")
    if err != nil {
      log.Fatal(err)
    }
    t.Execute(res, nil)
  })

  m.Post("/results", func(r *http.Request) string {
        text := r.FormValue("url")
        return text
    })
  m.Run()
}
