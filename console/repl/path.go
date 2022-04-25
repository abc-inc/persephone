// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repl

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PathComp returns directory entries matching the given path prefix. The path
// before the last path separator is used as base directory. If path is not
// absolute, then the path is resolved relatively to the working directory.
// If an I/O error occurs, no items are returned.
func PathComp(path string) []Item {
	if path == "" {
		path = "./"
	}

	idx := strings.LastIndex(path, "/")
	if idx == -1 || (!filepath.IsAbs(path) && !strings.HasPrefix(path, "./")) {
		path = "./" + path
		idx += 2
	}

	return findFromPrefix(path, idx)
}

// findFromPrefix returns files and directories matching the given path prefix.
// If path is not a directory, then the last path segment is used as a filter.
func findFromPrefix(path string, idx int) (is []Item) {
	var filter string
	if idx != len(path)-1 || !isDir(path) {
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

// fileDetails returns file size in bytes and modification time formatted as
// human-readable string.
func fileDetails(fi fs.FileInfo) string {
	return fmt.Sprintf("%8d %s", fi.Size(), fi.ModTime().Format(time.Stamp))
}

// isDir reports whether path is a directory.
func isDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}
