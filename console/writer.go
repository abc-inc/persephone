package console

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"text/tabwriter"

	"github.com/abc-inc/gutenfmt/formatter"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
)

func WriteTable(i interface{}) string {
	b := &strings.Builder{}
	tw := tabwriter.NewWriter(b, 4, 4, 1, ' ', 0)
	internal.Must(writeMapSlice(tw, reflect.ValueOf(i)))
	return b.String()
}

func WriteText(i interface{}, Sep, Delim string) string {
	f := FromStructSlice(Sep, Delim, i.([]graph.Result))
	s, err := f.Format(i)
	if err != nil {
		return ""
	}
	return s
}

func writeMapSlice(tw *tabwriter.Writer, v reflect.Value) (int, error) {
	f := FromStructSlice("\t", "\t\n", v.Interface().([]graph.Result))
	return formatter.FormatTab(tw, f, v.Interface())
}

func FromStructSlice(sep, delim string, typ []graph.Result) formatter.Formatter {
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
			b.WriteString(ToString(e.Values[0]))
			for pIdx := 1; pIdx < len(fs); pIdx++ {
				b.WriteString(sep)
				b.WriteString(ToString(e.Values[pIdx]))
			}
			b.WriteString(delim)
		}
		return b.String()[len(sep) : b.Len()-len(delim)], nil
	})
}

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	typ := reflect.TypeOf(i)
	switch typ.Kind() {
	case reflect.Array, reflect.Slice:
		s := fmt.Sprint(i)
		return s[1 : len(s)-1]
	case reflect.Chan:
		return typ.String()
	case reflect.Func:
		return funcName(reflect.ValueOf(i))
	case reflect.Ptr:
		return ToString(reflect.Indirect(reflect.ValueOf(i)).Interface())
	case reflect.String:
		return i.(string)
	default:
		return fmt.Sprint(i)
	}
}

// funcName returns the name of the function f points to.
func funcName(f reflect.Value) string {
	return runtime.FuncForPC(f.Pointer()).Name()
}
