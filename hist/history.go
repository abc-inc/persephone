package hist

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/internal"
	"github.com/pkg/errors"
)

type History struct {
	path    string
	entries []string
}

var history = History{}

func Get() *History {
	return &history
}

func Load(path string) *History {
	history = History{path: path}
	if histFile, err := os.Open(path); err == nil {
		s := bufio.NewScanner(histFile)
		for s.Scan() {
			history.entries = append(history.entries, s.Text())
		}
		internal.MustNoErr(histFile.Close())
	}
	return &history
}

func (h *History) Add(cyp string) {
	h.entries = append(h.entries, cyp)
	if len(h.entries) > 1000 {
		h.entries = h.entries[len(h.entries)-1000:]
	}
}

func (h History) Save() error {
	if err := os.MkdirAll(filepath.Dir(h.path), 0700); err != nil {
		format.Writeln(errors.WithMessage(err, "cannot write history file"))
		return nil
	}

	histFile, err := os.OpenFile(h.path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		format.Writeln(errors.WithMessagef(err, "cannot write history file"))
		return nil
	}

	defer func(histFile *os.File) {
		internal.MustNoErr(histFile.Close())
	}(histFile)

	_, err = h.WriteTo(histFile)
	return err
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
