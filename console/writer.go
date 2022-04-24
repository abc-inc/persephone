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
	"encoding/csv"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/gutenfmt/gfmt"
	"github.com/abc-inc/persephone/graph"
	"github.com/dustin/go-humanize/english"
	"github.com/fatih/color"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

// Query executes a statement and returns the result.
func Query(r graph.Request) error {
	log.Debug().Str("statement", r.Query).Fields(r.Params).Msg("Executing")

	if FormatName() == "raw" || FormatName() == "rawc" {
		return queryRaw(r)
	}
	return queryResult(r)
}

func queryRaw(req graph.Request) error {
	sp := NewSpinner()
	sp.Start()

	t := graph.NewTypedTemplate[map[string]interface{}](graph.GetConn())
	ms, sum, err := t.Query(r, graph.NewRawResultMapper())

	sp.Stop()
	if err == nil {
		Write(ms)
		writeSummary(len(ms), sum)
	}
	return err
}

func queryResult(req graph.Request) error {
	sp := NewSpinner()
	sp.Start()

	t := graph.NewTypedTemplate[graph.Result](graph.GetConn())
	rs, sum, err := t.Query(r, graph.NewResultMapper())

	sp.Stop()
	if err == nil {
		err = WriteResult(rs)
		writeSummary(len(rs), sum)
	}
	return err
}

func Write(i interface{}) {
	if _, err := w.Write(i); err != nil {
		log.Fatal().Err(err).Send()
	}
	fmt.Println()
}

// WriteErr formats the error message and writes it to stdout.
func WriteErr(err error) {
	if err == nil {
		return
	}

	msg := err.Error()
	r := []rune(msg[0:1])
	r[0] = unicode.ToUpper(r[0])
	color.Red(string(r) + msg[1:])
}

func WriteResult(rs []graph.Result) error {
	result, err := collectProps(rs)
	if err != nil {
		return err
	}

	var txt string
	switch wr := w.(type) {
	case gfmt.Tab:
		txt, err = writeTable(result)
	case gfmt.Text:
		txt, err = writeText(result, wr.Sep, wr.Delim)
	default: // json and yaml
		ms := []map[string]interface{}{}
		for _, r := range result {
			m := map[string]interface{}{}
			ms = append(ms, m)
			for i, k := range r.Keys {
				m[k] = r.Values[i]
				if props, ok := m[k].(map[string]any); ok {
					delete(props, graph.Label)
					delete(props, graph.Labels)
					delete(props, graph.Type)
				}
			}
		}
		Write(ms)
	}

	if txt != "" {
		Write(txt)
	}
	return err
}

// writeSummary returns a summary message of the executed query.
func writeSummary(n int, sum neo4j.ResultSummary) {
	const sumMsg = "%d %s, ready to start consuming query after %s, results consumed after another %s\n"
	log.Info().Msgf(sumMsg,
		n, english.PluralWord(n, "row", "rows"),
		sum.ResultAvailableAfter(), sum.ResultConsumedAfter())
}

func collectProps(rs []graph.Result) ([]graph.Result, error) {
	tm := GetTmplMgr()
	result := []graph.Result{}
	for i, r := range rs {
		result = append(result, graph.Result{})
		for j, v := range r.Values {
			props, ok := v.(map[string]interface{})
			if !ok {
				result[i].Add(r.Keys[j], v)
				continue
			}

			if l, ok := props[graph.Label]; ok && tm.Get(l.(string)) != nil {
				str, err := apply(tm.Get(l.(string)), props)
				if err != nil {
					return result, err
				}
				result[i].Add(r.Keys[j], str)
			} else if t, ok := props[graph.Type]; ok && tm.Get(t.(string)) != nil {
				str, err := apply(tm.Get(t.(string)), props)
				if err != nil {
					return result, err
				}
				result[i].Add(r.Keys[j], str)
			} else if strings.HasPrefix(FormatName(), "json") || strings.HasPrefix(FormatName(), "yaml") {
				result[i].Add(r.Keys[j], props)
			} else {
				str, err := toJSON(props)
				if err != nil {
					return result, err
				}
				result[i].Add(r.Keys[j], str)
			}
		}
	}
	return result, nil
}

func apply(t *template.Template, props map[string]interface{}) (string, error) {
	b := &strings.Builder{}
	if err := t.Execute(b, props); err != nil {
		return "", err
	}
	return strings.TrimSuffix(b.String(), "\n"), nil
}

func toJSON(props map[string]interface{}) (txt string, err error) {
	bs, err := json.Marshal(props)
	if err != nil {
		return
	}
	txt = string(bs)
	if FormatName() == "csv" {
		b := &strings.Builder{}
		w := csv.NewWriter(b)
		_ = w.Write([]string{txt})
		w.Flush()
		txt = strings.TrimSuffix(b.String(), "\n")
	}
	return
}

func writeTable(rs []graph.Result) (string, error) {
	b := &strings.Builder{}
	tw := tabwriter.NewWriter(b, 4, 4, 1, ' ', 0)
	_, err := writeMapSlice(tw, rs)
	return b.String(), err
}

func writeText(rs []graph.Result, sep, delim string) (string, error) {
	f := fromStructSlice(rs, sep, delim)
	return f.Format(rs)
}

func writeMapSlice(tw *tabwriter.Writer, rs []graph.Result) (int, error) {
	f := fromStructSlice(rs, "\t", "\t\n")
	return formatter.FormatTab(tw, f, rs)
}

func fromStructSlice(rs []graph.Result, sep, delim string) formatter.Formatter {
	if len(rs) == 0 {
		return formatter.NoopFormatter()
	}

	fs := rs[0].Keys

	return formatter.Func(func(i interface{}) (string, error) {
		rs := i.([]graph.Result)
		b := &strings.Builder{}
		for _, f := range fs {
			b.WriteString(sep)
			b.WriteString(f)
		}
		b.WriteString(delim)

		for idx := 0; idx < len(rs); idx++ {
			e := rs[idx]
			b.WriteString(toString(e.Values[0]))
			for pIdx := 1; pIdx < len(fs); pIdx++ {
				b.WriteString(sep)
				b.WriteString(toString(e.Values[pIdx]))
			}
			b.WriteString(delim)
		}
		return b.String()[len(sep) : b.Len()-len(delim)], nil
	})
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}
	typ := reflect.TypeOf(i)
	switch typ.Kind() {
	case reflect.Array, reflect.Slice:
		s := fmt.Sprint(i)
		return s[1 : len(s)-1]
	case reflect.Ptr:
		return toString(reflect.Indirect(reflect.ValueOf(i)).Interface())
	case reflect.String:
		return i.(string)
	default:
		return fmt.Sprint(i)
	}
}
