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

package console

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/abc-inc/persephone/internal"
	"github.com/rs/zerolog/log"
)

// TmplExt is the filename extension for templates.
const TmplExt = ".tmpl"

var tmplMgr *TmplMgr
var TmplDir = filepath.Join(internal.Must(os.UserConfigDir()), "persephone", "templates")

// NamedTemplate holds metadata of a template.
type NamedTemplate struct {
	Name       string `json:"Name"`
	Tmpl       string `json:"Template"`
	Persistent bool   `json:"Persistent"`
}

// String returns the template name.
func (t NamedTemplate) String() string {
	return t.Name
}

// TmplMgr loads and saves templates.
type TmplMgr struct {
	TmplsByPath map[string]*template.Template
}

// NewTmplMgr creates a new template manager.
// It manages two transient and persistent templates.
// Persistent templates are stored in the filesystem at an absolute path.
func NewTmplMgr() *TmplMgr {
	tmplMgr = &TmplMgr{TmplsByPath: make(map[string]*template.Template)}
	return tmplMgr
}

// GetTmplMgr returns the default template manager or creates a new one.
func GetTmplMgr() *TmplMgr {
	if tmplMgr != nil {
		return tmplMgr
	}
	return NewTmplMgr()
}

// Load parses all template in TmplDir.
// Already loaded templates with the same name are replaced.
func (tm *TmplMgr) Load() error {
	fs := internal.Must(filepath.Glob(filepath.Join(TmplDir, "*"+TmplExt)))
	for _, f := range fs {
		log.Info().Str("name", filepath.Base(f)).Msg("Loading template")
		if t, err := template.ParseFiles(f); err != nil {
			return err
		} else {
			tm.TmplsByPath[filepath.Base(f)] = t
		}
	}
	return nil
}

// Get returns the parsed Template for the given path.
func (tm *TmplMgr) Get(path string) (t *template.Template) {
	if !strings.HasSuffix(path, TmplExt) {
		path += TmplExt
	}
	var ok bool
	if t, ok = tm.TmplsByPath[path]; !ok {
		t, _ = tm.TmplsByPath[filepath.Join(TmplDir, path)]
	}
	return
}

// Set creates or replaces a template with the given text.
// path can be an absolute path or a basename without slashes.
// If it is an absolute path, the template is saved in that file.
func (tm *TmplMgr) Set(path, text string) (t *template.Template, err error) {
	name := filepath.Base(path)
	t = template.New(name)
	if t, err = t.Parse(text); err != nil {
		return nil, err
	}
	tm.TmplsByPath[name] = t
	if isPersistent(path) {
		if err = os.MkdirAll(filepath.Dir(path), 0750); err != nil {
			log.Err(err).Msg("Cannot save template")
			return
		}
		log.Info().Str("name", path).Msg("Saving template")
		err = os.WriteFile(path, []byte(text), 0600)
	}
	return
}

// isPersistent returns whether a template is stored in the filesystem.
func isPersistent(name string) bool {
	return filepath.IsAbs(name)
}