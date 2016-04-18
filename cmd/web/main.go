package main

import (
  "pr_helper"

  "log"
  "strings"
  "strconv"
  "fmt"

  "net/http"
  "html/template"
  "github.com/go-martini/martini"
)

func showPr(url string) string {
  ret := "<p>Analyzing pr " + url
  slices := strings.Split(url, "/")
  fmt.Println(slices)
  num, err := strconv.Atoi(slices[len(slices)-1])
  if err != nil {
    ret += "<p>Unable to find number " + slices[len(slices)-1]
    return ret
  }

  repo := pr_helper.RepositoryFromSettings()
  pr, err := repo.GetPR(num)
  if err != nil {
    ret += "<p>Unable to find pr " + strconv.Itoa(num)
    return ret
  }
  ret += "<p>"
  ret += pr.Topic()

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
      url := r.FormValue("url")
      return showPr(url)
    })
  m.Run()
}
