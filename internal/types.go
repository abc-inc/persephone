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

package internal

import (
	"strconv"
	"strings"
)

// Parse attempts to interpret the string as boolean or integer and returns it.
// Otherwise, it returns the value itself, regardless of its type.
func Parse(s string) any {
	if lc := strings.ToLower(s); lc == "true" || lc == "false" {
		return lc == "true"
	} else if val, err := strconv.ParseInt(s, 10, 32); err == nil {
		return val
	}
	return s
}

// ReSlice creates a new slice of T and inserts all elements from the original.
// It is assumed that all elements are of type T.
func ReSlice[T any](es []any) []T {
	ts := make([]T, len(es))
	for i, e := range es {
		ts[i] = e.(T)
	}
	return ts
}
