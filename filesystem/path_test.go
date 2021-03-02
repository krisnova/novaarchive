package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHomeTildaHappy(t *testing.T) {
	envHome := os.Getenv("HOME")
	expected := filepath.Join(envHome, ".meeps")
	path := NewPath("~/.meeps")
	if path.absolute != expected {
		t.Errorf("actual %s != expected %s", path.absolute, expected)
	}
}
