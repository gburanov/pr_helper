package main

import (
  "github.com/fatih/color"
)

func display(authors map[Author]int) {
  green := color.New(color.FgGreen)
  for author, count := range authors {
    green.Println(author.asStr(), "[", count, "]")
  }
}
