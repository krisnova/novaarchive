package backends

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kris-nova/novaarchive/filesystem"
)

const (
	DynamicDirectoryMode = 0655
	DynamicDirectoryType = "DynamicDirectory"
)

type DynamicDirectory struct {
	directory *filesystem.Directory
}

func NewDynamicDirectoryBackend(dir *filesystem.Directory) *DynamicDirectory {
	return &DynamicDirectory{
		directory: dir,
	}
}

func (b *DynamicDirectory) Get(key string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filepath.Join(b.directory.String(), key))
	if err != nil {
		return "", fmt.Errorf("unable to find key: %v", err)
	}
	if len(fileBytes) < 1 {
		return "", nil // Do not return an error here as we want to support empty values
	}
	return string(fileBytes), nil
}

func (b *DynamicDirectory) Set(key string, value string) error {
	return ioutil.WriteFile(filepath.Join(b.directory.String(), key), []byte(value), DynamicDirectoryMode)
}

func (b *DynamicDirectory) Type() string {
	return DynamicDirectoryType
}
