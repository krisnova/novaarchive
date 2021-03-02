package filesystem

import (
	"fmt"
	"io/ioutil"
)

func TempDir() (*Directory, error) {
	created, err := ioutil.TempDir("", "nova-")
	if err != nil {
		return nil, fmt.Errorf("unable to create temporary directory: %v", err)
	}
	return NewDirectory(created), nil
}
