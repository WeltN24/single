// package single provides a mechanism to ensure, that only one instance of a program is running

package single

import (
	"errors"
	"os"
	"time"
)

var (
	// ErrAlreadyRunning
	ErrAlreadyRunning = errors.New("the program is already running")
	//
	Lockfile string
)

// Single represents the name and the open file descriptor
type Single struct {
	name   string
	file   *os.File
	Locked bool
}

// New creates a Single instance
func New(name string) *Single {
	return &Single{name: name, Locked: false}
}

// Wait until the lock is released
func (s *Single) Wait() {
	locked := true
	for locked {
		time.Sleep(time.Millisecond)

		err := s.Lock()
		locked = err != nil

		if err == nil {
			s.Unlock()
		}
	}
}
