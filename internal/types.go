package internal

import "strconv"

func Parse(s string) (val interface{}) {
	var err error
	if val, err = strconv.ParseBool(s); err == nil {
	} else if val, err = strconv.ParseInt(s, 10, 32); err == nil {
	} else {
		val = s
	}
	return
}

func Reslice[T any](es []interface{}) (ts []T) {
	for _, e := range es {
		ts = append(ts, e.(T))
	}
	return ts
}
