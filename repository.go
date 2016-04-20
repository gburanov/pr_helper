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
		ret = append(ret, *pr)
	}
	return ret
}

func (repo *Repository) PRs() []PR {
	query := "is:open repo:" + repo.Organization + "/" + repo.Project + " label:" + GetSettings().Label
	return repo.listPRsbyQuery(query)
}

func (repo *Repository) GetPR(number int) (*PR, error) {
	pr := PR{Repository: repo, Number: number}
	return &pr, pr.validate()
}

func (repo *Repository) LocalPath() string {
	return GetSettings().RepositoryPath + "/" + repo.Organization + "/" + repo.Project
}

func (repo *Repository) RootPath() string {
	return GetSettings().RepositoryPath + "/" + repo.Organization
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
		if err != nil { repo.Create(cb) }
	}
	// Update repo
	command := exec.Command("git", "pull")
	command.Dir = repo.LocalPath()
	return command.Run()
}

func (repo *Repository) Create(cb Callback) {
	cb("Creating %s...", repo.LocalPath())
	err := exec.Command("mkdir", "-p", repo.RootPath()).Run()
	if err != nil {
		log.Fatal(err)
	}
	// Remove subdirectory
	command := exec.Command("rm", "-rf", repo.Project)
	command.Dir = repo.RootPath()
	err = command.Run()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("https://%s@github.com/%s/%s.git",
		GetSettings().AuthToken, repo.Organization, repo.Project)
	fmt.Println("git clone %s",path)
	command = exec.Command("git", "clone", path)
	command.Dir = repo.RootPath()
	err = command.Run()
	if err != nil {
		log.Fatal(err)
	}
}
