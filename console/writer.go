package console

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/dustin/go-humanize/english"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
)

func WriteResult(rs []graph.Result, sum neo4j.ResultSummary) error {
	result, err := process(rs)
	if err != nil {
		return err
	}

	if FormatName() == "table" {
		Writeln(WriteTable(result))
	} else if FormatName() == "csv" {
		Writeln(WriteText(result, ",", "\n"))
	} else if FormatName() == "text" {
		Writeln(WriteText(result, ";", "\n"))
	} else if FormatName() == "tsv" {
		Writeln(WriteText(result, "\t", "\t\n"))
	} else if strings.HasPrefix(FormatName(), "json") || strings.HasPrefix(FormatName(), "yaml") {
		ms := []map[string]interface{}{}
		for _, r := range result {
			m := map[string]interface{}{}
			ms = append(ms, m)
			for i, k := range r.Keys {
				m[k] = r.Values[i]
				if props, ok := m[k].(map[string]any); ok {
					delete(props, "@label")
					delete(props, "@labels")
				}
			}
		}
		Writeln(ms)
	} else {
		Writeln(result)
	}

	return nil
}

func WriteSummary(n int, sum neo4j.ResultSummary) {
	const sumMsg = "%d %s, ready to start consuming query after %s, results consumed after another %s\n"
	log.Info().Msgf(sumMsg,
		n, english.PluralWord(n, "row", "rows"),
		sum.ResultAvailableAfter(), sum.ResultConsumedAfter())
}

func process(rs []graph.Result) ([]graph.Result, error) {
	result := []graph.Result{}
	for i, r := range rs {
		result = append(result, graph.Result{})
		for j, v := range r.Values {
			if props, ok := v.(map[string]interface{}); ok {
				if l, ok := props["@label"]; ok && Tmpls[l.(string)+".tmpl"] != nil {
					tmpl := Tmpls[l.(string)+".tmpl"]
					b := &strings.Builder{}
					if err := tmpl.Execute(b, props); err != nil {
						return result, err
					}
					result[i].Add(r.Keys[j], strings.TrimSuffix(b.String(), "\n"))
				} else if l, ok := props["@type"]; ok && Tmpls[l.(string)+".tmpl"] != nil {
					tmpl := Tmpls[l.(string)+".tmpl"]
					b := &strings.Builder{}
					if err := tmpl.Execute(b, props); err != nil {
						return result, err
					}
					result[i].Add(r.Keys[j], strings.TrimSuffix(b.String(), "\n"))
				} else if FormatName() == "json" || FormatName() == "jsonc" || FormatName() == "yaml" || FormatName() == "yamlc" {
					result[i].Add(r.Keys[j], props)
				} else {
					bs, err := json.Marshal(props)
					if err != nil {
						return result, err
					}
					if FormatName() == "csv" {
						b := &strings.Builder{}
						w := csv.NewWriter(b)
						_ = w.Write([]string{string(bs)})
						w.Flush()
						result[i].Add(r.Keys[j], strings.TrimSuffix(b.String(), "\n"))
					} else {
						result[i].Add(r.Keys[j], string(bs))
					}
				}
			} else {
				result[i].Add(r.Keys[j], v)
			}
		}
	}
	return result, nil
}

func WriteTable(i interface{}) string {
	b := &strings.Builder{}
	tw := tabwriter.NewWriter(b, 4, 4, 1, ' ', 0)
	internal.Must(writeMapSlice(tw, reflect.ValueOf(i)))
	return b.String()
}

func WriteText(i interface{}, Sep, Delim string) string {
	f := fromStructSlice(Sep, Delim, i.([]graph.Result))
	s, err := f.Format(i)
	if err != nil {
		return ""
	}
	return s
}

func writeMapSlice(tw *tabwriter.Writer, v reflect.Value) (int, error) {
	f := fromStructSlice("\t", "\t\n", v.Interface().([]graph.Result))
	return formatter.FormatTab(tw, f, v.Interface())
}

func fromStructSlice(sep, delim string, typ []graph.Result) formatter.Formatter {
	if len(typ) == 0 {
		return formatter.NoopFormatter()
	}

	fs := typ[0].Keys

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
