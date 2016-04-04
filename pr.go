package pr_helper

import (
	"log"
	"strings"
)

type PR struct {
	Repository *Repository
	Number     int
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

func (pr *PR) Authors() *Authors {
	files, _, err := pr.Repository.Client.PullRequests.
		ListFiles(pr.Repository.Organization, pr.Repository.Project, pr.Number, nil)
	if err != nil {
		log.Fatal(err)
	}

	results := 0
	authors_channel := make(chan []Author)
	for _, file := range files {
		if strings.HasPrefix(*file.Filename, "phrase") {
			continue
		}
		results++
		go func(fileName string) {
			authors_channel <- fileAuthors(fileName)
		}(*file.Filename)
	}
	authors := []Author{}
	for i := 1; i <= results; i++ {
		authors = append(authors, <-authors_channel...)
	}
	return arrayToMap(authors)
}
