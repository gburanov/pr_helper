package main

import (
  "pr_helper"
  "fmt"
)

func showPr(url string) string {
  ret := "<p>Analyzing pr " + url
  manager := pr_helper.NewManager()
  pr, err := manager.GetPR(url)
  if err != nil {
    ret += err.Error()
    return ret
  }
  ret += "<p>"
  ret += pr.Topic()
  ret += showAuthors(*pr)

  return ret
}

func showLeftStats(authors *pr_helper.Authors) string {
  left, total := authors.GetLinesStat()
  percent := float32(left)/float32(total)

  ret := fmt.Sprintf("<p>%d out of %d lines unmntained", left, total)
  if (total > 100 && percent > 0.7) || (percent > 0.9) {
    ret += "<p>" + "WARNING! DEEP LEGACY"
  }
  return ret
}

func showAuthors(pr pr_helper.PR) string {
  ret := ""
  authors := pr.Authors()
  ret += showLeftStats(authors)
  for author, _ := range *authors {
    ret += "<p>" + author.AsStr()
  }
  return ret
}
