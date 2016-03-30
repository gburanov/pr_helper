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

func showLeftStats(authors map[Author]int) {
  total := 0
  left := 0
  for author, lines := range authors {
    if !author.AtWimdu() {
      left += lines
    }
    total += lines
  }
  percent := float32(left)/float32(total)
  yellow := color.New(color.FgYellow)
  yellow.Println(left, "out of", total,"lines unmntained")
  if (total > 100 && percent > 0.7) || (percent > 0.9) {
    yellow.Println("WARNING! DEEP LEGACY")
  }
}
