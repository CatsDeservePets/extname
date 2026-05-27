package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	progName = strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
	allExts  = flag.Bool("a", false, "print all extension segments (e.g. .tar.gz)")
	dropDot  = flag.Bool("d", false, "print the extension without a leading dot")
	zero     = flag.Bool("z", false, "end each output line with NUL, not newline")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-a] [-d] [-z] string [...]\n", progName)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}

	delim := "\n"
	if *zero {
		delim = "\x00"
	}

	for _, p := range flag.Args() {
		fmt.Print(Extension(p, *allExts, *dropDot), delim)
	}
}

// Extension returns the file extension of path.
// Unlike [filepath.Ext], Extension applies [filepath.Base] to path and
// ignores leading dots when determining the extension. The extension is the
// suffix beginning at the final dot, or at the first non-leading dot if all
// is true. If nodot is true, the leading dot is omitted from the extension.
// If no extension is found, Extension returns an empty string.
func Extension(path string, all, nodot bool) string {
	// Distinguish dotfiles from file extensions.
	name := strings.TrimLeft(filepath.Base(path), ".")
	if name == "" {
		return ""
	}
	var i int
	if all {
		i = strings.IndexByte(name, '.')
	} else {
		i = strings.LastIndexByte(name, '.')
	}
	if i < 0 {
		return ""
	}

	if nodot {
		i++
	}
	return name[i:]
}
