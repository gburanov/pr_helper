package pr_helper

import (
	"time"
)

type Mutex struct {
	c chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.c <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.c
}

func (m *Mutex) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
  select {
	case m.c <- struct{}{}:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}
