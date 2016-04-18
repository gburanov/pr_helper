package pr_helper

import (
	"github.com/google/go-github/github"
	"log"
	"fmt"
	"os/exec"
)

type Repository struct {
	Organization string
	Project      string
  Client       *github.Client
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

func (repo *Repository) LocalPath() string {
	return GetSettings().RepositoryPath + "/" + repo.Organization + "/" + repo.Project
}

func (repo *Repository) Init() {
	path := repo.LocalPath()
	exist, err := exists(path)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		repo.Create()
	} else {
		command := exec.Command("git", "status")
		command.Dir = repo.LocalPath()
		err := command.Run()
		if err != nil { repo.Create() }
	}
}

func (repo *Repository) Create() {
	fmt.Println("Creating", repo.LocalPath())
	err := exec.Command("mkdir", "-p", repo.LocalPath()).Run()
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("https://%s@github.com/%s/%s.git",
		GetSettings().AuthToken, repo.Organization, repo.Project)
	fmt.Println("Clonning", path)
	command := exec.Command("git", "clone", path, ".")
	command.Dir = repo.LocalPath()
	err = command.Run()
	if err != nil {
		log.Fatal(err)
	}
}
