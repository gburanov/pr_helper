package main

import (
	"github.com/fatih/color"
	"github.com/gburanov/pr_helper/lib"
)

func displayAuthors(authors *pr_helper.Authors) {
	green := color.New(color.FgGreen)
	for author, count := range *authors {
		green.Println(author.AsStr(), "[", count, "]")
	}
}

func displayPR(pr pr_helper.PR) {
	red := color.New(color.FgRed, color.Bold)
	red.Println(pr.Topic())
	red.Println(pr.Url())

	authors := pr.Authors()
	authors.Check()

	showLeftStats(authors)

	authors = pr_helper.FilterTop(5, authors)
	displayAuthors(authors)
}

func showLeftStats(authors *pr_helper.Authors) {
	left, total := authors.GetLinesStat()
	percent := float32(left) / float32(total)
	yellow := color.New(color.FgYellow)
	yellow.Println(left, "out of", total, "lines unmntained")
	if (total > 100 && percent > 0.7) || (percent > 0.9) {
		yellow.Println("WARNING! DEEP LEGACY")
	}
}
