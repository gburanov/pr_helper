package pr_helper

import (
	"strings"
	"errors"
	"strconv"
	"time"
	"github.com/google/go-github/github"
)

type Manager struct {
  Client       *github.Client
	Cb Callback
	M Mutex
}

func NewManager(cb Callback, m *Mutex) *Manager {
  return &Manager{
		Client: github.NewClient(Token()),
		Cb: cb,
		M: *m,
	}
}

func (m *Manager) GetRepository(organization string, project string) (*Repository, error) {
  repo := &Repository{
		Organization: organization,
		Project: project,
		Client: m.Client,
	}
	m.Cb("Trying to lock repository...")
	locked := m.M.TryLock(time.Second * 10)
	if locked == false {
		m.Cb("Failed to lock please try again later")
		return nil, errors.New("Locked repository")
	} else {
		m.Cb("Locked succesfully")
	}
  err := repo.Init(m.Cb)
	m.M.Unlock()
	fmt.Println("Unlocked back")
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
