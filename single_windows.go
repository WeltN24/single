// +build windows

package single

import (
	"fmt"
	"os"
	"path/filepath"
)

// Filename returns an absolute filename, appropriate for the operating system
func (s *Single) Filename() string {
	if len(Lockfile) > 0 {
		return Lockfile
	}
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s.lock", s.name))
}

// Lock tries to remove the lock file, if it exists.
// If the file is already open by another instance of the program,
// remove will fail and exit the program.
func (s *Single) Lock() error {

	if err := os.Remove(s.Filename()); err != nil && !os.IsNotExist(err) {
		return ErrAlreadyRunning
	}

	file, err := os.OpenFile(s.Filename(), os.O_EXCL|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	s.file = file

	return nil
}

// Unlock closes and removes the lockfile.
func (s *Single) Unlock() error {
	if err := s.file.Close(); err != nil {
		return err
	}
	if err := os.Remove(s.Filename()); err != nil {
		return err
	}

	return nil
}
