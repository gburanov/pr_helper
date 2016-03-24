package main

import "github.com/fatih/color"

func display(authors map[string]int) {
  green := color.New(color.FgGreen)
  for author, count := range authors {
    green.Println(author, "[", count, "]")
  }
}
