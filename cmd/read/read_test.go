package main

import (
	"testing"
	"manifest/manifest"
)

func TestRead(t *testing.T) {
	paths := []string{"/home/owner/"}
	manifest.ReadFiles(paths)
}
