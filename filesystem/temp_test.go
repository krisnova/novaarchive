package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTempDirHappy(t *testing.T) {
	dir, err := TempDir()
	defer dir.Remove()
	if err != nil {
		t.Errorf("unable to create temp dir: %v", err)
		t.FailNow()
	}
	// Ensure we can write to it
	err = ioutil.WriteFile(filepath.Join(dir.String(), "meeps"), []byte("SolarWinds123"), 0655)
	if err != nil {
		t.Errorf("unable to write file to temp dir %s: %v", dir.String(), err)
		t.FailNow()
	}
}

func TestDirEnsure(t *testing.T) {
	dir := NewDirectory(filepath.Join(os.TempDir(), "meeps"))
	defer dir.Remove()
	err := dir.Ensure(0655)
	if err != nil {
		t.Errorf("unable to ensure directory %s: %v", dir.String(), err)
		t.FailNow()
	}
}

func TestDirRemove(t *testing.T) {
	dir := NewDirectory(filepath.Join(os.TempDir(), "meeps"))
	err := dir.Ensure(0655)
	if err != nil {
		t.Errorf("unable to ensure directory %s: %v", dir.String(), err)
		t.FailNow()
	}
	err = dir.Remove()
	if err != nil {
		t.Errorf("unable to remove directory %s: %v", dir.String(), err)
		t.FailNow()
	}
}
