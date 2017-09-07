package main

import (
	"testing"
	"manifest/manifest"
)

func BenchmarkRead(b *testing.B) {
	paths := []string{"/home/owner/"}
	manifest.ReadFiles(paths)
}
