package pr_helper

import (
	"log"
	"strings"
)

type PR struct {
	Repository *Repository
	Number     int
}

func (pr *PR) validate() error {
	_, _, err := pr.Repository.Client.PullRequests.
		Get(pr.Repository.Organization, pr.Repository.Project, pr.Number)
	return err
}

func (pr *PR) Topic() string {
	pr_, _, err := pr.Repository.Client.PullRequests.
		Get(pr.Repository.Organization, pr.Repository.Project, pr.Number)
	if err != nil {
		log.Fatal(err)
	}
	return *pr_.Title
}

func (pr *PR) Url() string {
	pr_, _, err := pr.Repository.Client.PullRequests.
		Get(pr.Repository.Organization, pr.Repository.Project, pr.Number)
	if err != nil {
		log.Fatal(err)
	}
	return *pr_.HTMLURL
}

func (pr *PR) Stats() Stats {
	files, _, err := pr.Repository.Client.PullRequests.
		ListFiles(pr.Repository.Organization, pr.Repository.Project, pr.Number, nil)
	if err != nil {
		log.Fatal(err)
	}

	results := 0
	stats_channel := make(chan Stats)
	for _, file := range files {
		if strings.HasPrefix(*file.Filename, "phrase") {
			continue
		}
		results++
		go func(pr *PR, fileName string) {
			stats_channel <- fileStatistics(pr.Repository, fileName)
		}(pr, *file.Filename)
	}
	stats := Stats{}
	for i := 1; i <= results; i++ {
		stats = append(stats, <-stats_channel...)
	}
	return stats
	//return arrayToMap(stats.Authors())
}
