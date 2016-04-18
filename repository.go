package pr_helper

import (
	"github.com/google/go-github/github"
	"log"
	"net/http"
)

type Repository struct {
	Organization string
	Project      string
}

func (repo *Repository) listPRsbyQuery(query string) []PR {
	prs, _, err := repo.Client.Search.Issues(query, &github.SearchOptions{})
	if err != nil {
		log.Fatal(err)
	}
	ret := []PR{}
	for _, github_pr := range prs.Issues {
		pr, err := repo.GetPR(*github_pr.Number)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, pr)
	}
	return ret
}

func (repo *Repository) MyPRs() []PR {
	query := "is:open repo:" + repo.Organization + "/" + repo.Project + " label:" + GetSettings().Label + " author:gburanov"
	return repo.listPRsbyQuery(query)
}

func (repo *Repository) PRs() []PR {
	query := "is:open repo:" + repo.Organization + "/" + repo.Project + " label:" + GetSettings().Label
	return repo.listPRsbyQuery(query)
}

func (repo *Repository) GetPR(number int) (PR, error) {
	pr := PR{Repository: repo, Number: number}
	return pr, pr.validate()
}
