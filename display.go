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
  for author, _ := range authors {
    if !author.available() {
      left += 1
    }
    total += 1
  }
  percent := float32(left)/float32(total)
  yellow := color.New(color.FgYellow)
  yellow.Println(left, "out of", total,"left")
  if (total > 10 && percent > 0.33) {
    yellow.Println("WARNING! DEEP LEGACY")
  }
}
