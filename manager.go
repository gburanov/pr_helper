package pr_helper

import (
	"github.com/google/go-github/github"
)

type Manager struct {
  Client       *github.Client
}

func NewManager() *Manager {
  return &Manager{Client: github.NewClient(Token())}
}

func (m *Manager) GetRepository(organization string, project string) *Repository {
  repo := new(Repository)
  repo.Organization = organization
  repo.Project = project
  repo.Client = m.Client
  repo.Init()
  return repo
}
