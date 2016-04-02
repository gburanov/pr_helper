package pr_helper

import (
	"github.com/google/go-github/github"
	"log"
	"net/http"
)

type Repository struct {
	Organization string
	Project      string
	Client       *github.Client
}

func NewRepository(organization string, project string, auth_token *http.Client) *Repository {
	repo := new(Repository)
	repo.Organization = organization
	repo.Project = project
	repo.Client = github.NewClient(auth_token)
	return repo
}

func (repo *Repository) listPRsbyQuery(query string) []PR {
	prs, _, err := repo.Client.Search.Issues(query, &github.SearchOptions{})
	if err != nil {
		log.Fatal(err)
	}
	ret := []PR{}
	for _, pr := range prs.Issues {
		ret = append(ret, *repo.GetPR(*pr.Number))
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

func (repo *Repository) GetPR(number int) *PR {
	pr := new(PR)
	pr.Repository = repo
	pr.Number = number
	return pr
}
