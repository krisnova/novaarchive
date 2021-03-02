package metrics

import (
	"testing"

	"github.com/kris-nova/novaarchive/filesystem"
)

func TestGetHappy(t *testing.T) {
	dir, err := filesystem.TempDir()
	if err != nil {
		t.Errorf("unable to create temp dir: %v", err)
		t.FailNow()
	}
	resource, err := NewDynamicDirectoryResource(dir)
	if err != nil {
		t.Errorf("unable to get new dynamic directory: %v", err)
		t.FailNow()
	}
	expected := "novaPants"
	resource.Set("myKey", expected)
	actual, err := resource.Get("myKey")
	if err != nil {
		t.Errorf("unable to get key: %v", err)
		t.FailNow()
	}
	if actual.Value != expected {
		t.Errorf("actual (%s) != expected (%s)", actual.Value, expected)
	}
}

func TestGetSad(t *testing.T) {
	dir, err := filesystem.TempDir()
	if err != nil {
		t.Errorf("unable to create temp dir: %v", err)
		t.FailNow()
	}
	resource, err := NewDynamicDirectoryResource(dir)
	if err != nil {
		t.Errorf("unable to get new dynamic directory: %v", err)
		t.FailNow()
	}
	// Do not call Set()
	_, err = resource.Get("myKey")
	if err == nil {
		t.Errorf("expected error getting bogus key 'myKey' in dir: %s", dir.String())
	}
}
