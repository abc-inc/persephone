package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/abc-inc/persephone/comp"
)

func PathCmpl(path string) comp.Result {
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

func findFromPrefix(path string, lastSepIdx int) (r comp.Result) {
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
				det = fmt.Sprintf("%8d %s", fi.Size(), fi.ModTime().Format(time.Stamp))
			}
		}
		r.Items = append(r.Items, comp.Item{View: de.Name(), Content: det})
	}
	return
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}