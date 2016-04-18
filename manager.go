package pr_helper

import (
	"net/http"

	"github.com/google/go-github/github"
)

type Manager struct {
  Client       *github.Client
}

func NewManager(auth_token *http.Client) *Manager {
  return &Manager{Client: github.NewClient(auth_token)}
}

func (m *Manager) GetRepository(organization string, project string) *Repository {
  repo := new(Repository)
  repo.Organization = organization
  repo.Project = project
  repo.Client = m.Client
  repo.Init()
  return repo
}
