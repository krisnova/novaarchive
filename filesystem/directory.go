package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// Path is used to represent a directory
// on the system.
//
// By design no information about the
// current context will be associated here.
type Directory struct {
	given    string
	absolute string
}

// NewDirectory represents a directory on the
// filesystem.
func NewDirectory(given string) *Directory {
	// we intentionally ignore errors
	// because we want this fail in the
	// implementation code.
	expanded := Expand(given)
	absolute, err := filepath.Abs(expanded)
	if err != nil {
		absolute = "" // <-- Ensure empty string
	}
	d := &Directory{
		given:    given,
		absolute: absolute,
	}
	return d
}

func (d *Directory) String() string {
	return d.absolute
}

// Ensure will attempt to ensure a directory exists
func (d *Directory) Ensure(mode os.FileMode) error {
	// check if absolute is valid
	if d.absolute == "" {
		return fmt.Errorf("invalid absolute path")
	}
	// check if dir is actually a file
	_, err := os.Stat(d.absolute)
	if os.IsExist(err) {
		// File exists
		return fmt.Errorf("file %s exists, unable to create directory", d.absolute)
	}
	if os.IsNotExist(err) {
		// Nothing exists
		return os.MkdirAll(d.absolute, mode)
	}
	return nil
}

// Remove will attempt to remove a directory and
// it's contents
func (d *Directory) Remove() error {
	return os.RemoveAll(d.absolute)
}
