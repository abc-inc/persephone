package fuzzy

import (
	"os"
	"strings"
)

func Search[T any](items []T, query string, keyEx func(T) string) (res []T) {
	if query == "" {
		return items
	}
	queryHasSlashes := strings.ContainsRune(query, '/')
	query = strings.ReplaceAll(query, " ", "")
	return filter(items, query, queryHasSlashes, keyEx)
}

func queryIsLastPathSegment(str, query string) bool {
	if len(str) > len(query) && str[len(str)-len(query)-1] == os.PathSeparator {
		return strings.LastIndex(str, query) == len(str)-len(query)
	}
	return false
}
