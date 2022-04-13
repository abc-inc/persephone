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
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/abc-inc/persephone/internal"
	"github.com/pkg/errors"
)

type History struct {
	Max     int
	entries []string
}

var _ io.WriterTo = (*History)(nil)
var hist *History

func NewHistory() *History {
	hist = &History{Max: 1000}
	return hist
}

func GetHistory() *History {
	if hist != nil {
		return hist
	}
	return NewHistory()
}

func (h *History) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	sc := bufio.NewScanner(f)
	h.Clear()
	for sc.Scan() {
		h.entries = append(h.entries, sc.Text())
	}
	return nil
}

func (h *History) Add(e string) {
	h.entries = append(h.entries, e)
	if len(h.entries) > h.Max {
		h.entries = h.entries[len(h.entries)-h.Max:]
	}
}

func (h History) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return errors.WithMessage(err, "cannot write history file")
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return errors.WithMessage(err, "cannot write history file")
	}

	defer func() { _ = f.Close() }()
	return internal.Second(h.WriteTo(f))
}

func (h History) WriteTo(w io.Writer) (n int64, err error) {
	for _, e := range h.entries {
		cnt, _ := io.WriteString(w, strings.ReplaceAll(e, "\n", " ")+"\n")
		n += int64(cnt)
	}
	return n, err
}

func (h *History) Clear() {
	h.entries = nil
}

func (h History) Entries() []string {
	return h.entries
}
