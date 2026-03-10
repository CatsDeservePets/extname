package main

import (
	"testing"
)

func TestExtension(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		all, nodot bool
		want       string
	}{
		{"empty", "", false, false, ""},
		{"no extension", "file", false, false, ""},

		{"single extension", "file.txt", false, false, ".txt"},
		{"single extension nodot", "file.txt", false, true, "txt"},

		{"multi-part extension", "archive.tar.gz", false, false, ".gz"},
		{"multi-part extension all", "archive.tar.gz", true, false, ".tar.gz"},
		{"multi-part extension nodot", "archive.tar.gz", false, true, "gz"},
		{"multi-part extension all nodot", "archive.tar.gz", true, true, "tar.gz"},

		{"dotfile", ".env", false, false, ""},
		{"dotfile with extension", ".env.local", false, false, ".local"},
		{"dotfile with extension nodot", ".env.local", false, true, "local"},

		{"trailing dot", "file.", false, false, "."},
		{"trailing dot nodot", "file.", false, true, ""},

		{"only dots", "...", false, false, ""},
		{"path with trailing slash", "/tmp/archive.tar.gz/", false, false, ".gz"},
		{"dots in directory", "dir.with.dots/file.txt", false, false, ".txt"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Extension(test.path, test.all, test.nodot)
			if got != test.want {
				t.Errorf("Extension(%q, %t, %t) = %q; want %q", test.path, test.all, test.nodot, got, test.want)
			}
		})
	}
}
