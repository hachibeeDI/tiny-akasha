package helper

import (
	"os"
	"path/filepath"
)

var (
	DirName string
)

func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	DirName = dir
}
