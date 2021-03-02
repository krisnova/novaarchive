package filesystem

import (
	"path/filepath"
	"strings"
)

// Path is used to represent a filepath
// on the system.
//
// By design no information about the
// current context will be associated here.
type Path struct {
	given    string
	absolute string
}

func NewPath(given string) *Path {
	// we intentionally ignore errors
	// because we want this fail in the
	// implementation code.
	expanded := Expand(given)
	absolute, err := filepath.Abs(expanded)
	if err != nil {
		absolute = "" // <-- Ensure empty string
	}
	p := &Path{
		given:    given,
		absolute: absolute,
	}
	return p
}

// Expand will expand common parlance
// literally for known UNIX parlance
//
// - Home
// - TODO .
// - TODO ..
func Expand(given string) string {
	if strings.Contains(given, "~") {
		return strings.Replace(given, "~", Home(), 1)
	}
	return given
}
