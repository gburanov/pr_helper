package main

import (
  "github.com/fatih/color"
  "github.com/google/go-github/github"
)

func display(authors map[string]int) {
  green := color.New(color.FgGreen)
  for author, count := range authors {
    green.Println(author, "[", count, "]")
  }
}

func displayPR(client *github.Client, num int) {
  authors := processPr(client, num)
  authors = filterTop(5, authors)
  display(authors)
}
