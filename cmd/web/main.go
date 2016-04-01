package main

import (
	"io"
	"net/http"

  "pr_helper"
)

func hello(w http.ResponseWriter, r *http.Request) {
  repo := pr_helper.NewRepository(pr_helper.GetSettings().Organization, pr_helper.GetSettings().Project, pr_helper.Token())
  for _, pr := range repo.PRs() {
		io.WriteString(w, "<p>")
    io.WriteString(w, pr.ShowInfo())
  }
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
