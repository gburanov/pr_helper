package main

import (
	"github.com/gburanov/pr_helper"
)

func showPr(url string, cb pr_helper.Callback, m *pr_helper.Mutex) {
	cb("Analyzing pr %s", url)
	manager := pr_helper.NewManager(cb, m)
	pr, err := manager.GetPR(url)
	if err != nil {
		cb(err.Error())
		return
	}
	cb(pr.Topic())
	showAuthors(*pr, cb)
}

func showLeftStats(authors *pr_helper.Authors, cb pr_helper.Callback) {
	left, total := authors.GetLinesStat()
	percent := float32(left) / float32(total)
	if left == 0 {
		// does not make sense
		return
	}

	cb("%d out of %d lines unmntained", left, total)
	if (total > 100 && percent > 0.7) || (percent > 0.9) {
		cb("WARNING! DEEP LEGACY")
	}
}

func showAuthors(pr pr_helper.PR, cb pr_helper.Callback) {
	stats := pr.Stats()
	cb("Total files %d", stats.FilesCount)
	cb("Total lines %d", stats.Lines())
	cb("Average time %s", stats.AverageTime().String())
	cb("Earliest time %s", stats.EarliestTime().String())
	authors := pr_helper.CreateAuthors(stats)
	showLeftStats(authors, cb)
	for author, lines := range *pr_helper.FilterTop(5, authors) {
		cb("%s (%d lines changed)", author.AsStr(), lines)
	}
}
