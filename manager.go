package pr_helper

import (
	"strings"
	"errors"
	"strconv"
	"github.com/google/go-github/github"
)

type Manager struct {
  Client       *github.Client
}

func NewManager() *Manager {
  return &Manager{Client: github.NewClient(Token())}
}

func (m *Manager) GetRepository(organization string, project string) (*Repository, error) {
  repo := new(Repository)
  repo.Organization = organization
  repo.Project = project
  repo.Client = m.Client
  err := repo.Init()
	if err != nil {
		return nil, err
	}
  return repo, nil
}

func (m *Manager) GetPR(url string) (*PR, error) {
	slices := strings.Split(url, "/")
	if len(slices) < 5 {
		return nil, errors.New("Bad URL " + url)
	}
	num, err := strconv.Atoi(slices[len(slices)-1])
	if err != nil {
		return nil, err
	}
	project := slices[len(slices)-3]
	organization := slices[len(slices)-4]
	repo, err := m.GetRepository(organization, project)
	if err != nil {
		return nil, err
	}
	pr, err := repo.GetPR(num)
	if err != nil {
		return nil, err
	}
	return pr, nil
}
