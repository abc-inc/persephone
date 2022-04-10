package repl

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/abc-inc/persephone/fuzzy"
)

func FileCompFor(dir string, infoFunc func(info os.FileInfo) string) CompFunc {
	if infoFunc == nil {
		infoFunc = func(info os.FileInfo) string { return "" }
	}

	return func(name string) (its []Item) {
		des, err := os.ReadDir(dir)
		if err != nil {
			return nil
		}
		des = fuzzy.Search(des, name, func(e os.DirEntry) string {
			return e.Name()
		})

		for _, de := range des {
			if fi, err := de.Info(); err == nil {
				its = append(its, Item{View: de.Name(), Content: infoFunc(fi)})
			}
		}
		return
	}
}

func PathComp(path string) []Item {
	if path == "" {
		path = "./"
	}

	lastSepIdx := strings.LastIndex(path, "/")
	if lastSepIdx == -1 || (!filepath.IsAbs(path) && !strings.HasPrefix(path, "./")) {
		path = "./" + path
		lastSepIdx += 2
	}

	return findFromPrefix(path, lastSepIdx)
}

func findFromPrefix(path string, lastSepIdx int) (is []Item) {
	var filter string
	if lastSepIdx != len(path)-1 || !isDir(path) {
		// path does not end with /
		// it might be directory or incomplete name - try parent directory
		filter = filepath.Base(path)
		path = filepath.Dir(path)
	}

	des, err := os.ReadDir(path)
	if err != nil {
		// if an I/O error occurs, no completion is provided
		return
	}

	for _, de := range des {
		if !strings.HasPrefix(de.Name(), filter) {
			continue
		}

		var det string
		if !de.IsDir() {
			if fi, err := de.Info(); err == nil {
				det = fileDetails(fi)
			}
		}
		is = append(is, Item{View: de.Name(), Content: det})
	}
	return
}

func fileDetails(fi fs.FileInfo) string {
	return fmt.Sprintf("%8d %s", fi.Size(), fi.ModTime().Format(time.Stamp))
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}
