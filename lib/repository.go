package pr_helper

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/github"
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
		ret = append(ret, *pr)
	}
	return ret
}

func (repo *Repository) PRs() []PR {
	query := "is:open repo:" + repo.Organization + "/" + repo.Project + " label:" + getSettings().Label
	return repo.listPRsbyQuery(query)
}

func (repo *Repository) GetPR(number int) (*PR, error) {
	pr := PR{Repository: repo, Number: number}
	return &pr, pr.validate()
}

func (repo *Repository) LocalPath() string {
	return getSettings().RepositoryPath + "/" + repo.Organization + "/" + repo.Project
}

func (repo *Repository) RootPath() string {
	return getSettings().RepositoryPath + "/" + repo.Organization
}

func (repo *Repository) Init(cb Callback) error {
	path := repo.LocalPath()
	exist, err := exists(path)
	if err != nil {
		return err
	}
	if !exist {
		repo.Create(cb)
	} else {
		command := exec.Command("git", "status")
		command.Dir = repo.LocalPath()
		err := command.Run()
		if err != nil {
			repo.Create(cb)
		}
	}
	// Update repo
	command := exec.Command("git", "pull")
	command.Dir = repo.LocalPath()
	return command.Run()
}

func (repo *Repository) Create(cb Callback) {
	cb("Creating %s...", repo.LocalPath())
	err := repo.ExecuteCommandInDir("", "mkdir", "-p", repo.RootPath())
	if err != nil {
		log.Fatal(err)
	}
	// Remove subdirectory
	err = repo.ExecuteCommandInDir(repo.RootPath(), "rm", "-rf", repo.Project)
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("https://%s@github.com/%s/%s.git",
		getSettings().AuthToken, repo.Organization, repo.Project)
	err = repo.ExecuteCommandInDir(repo.RootPath(), "git", "clone", path)
	if err != nil {
		log.Fatal(err)
	}
}
