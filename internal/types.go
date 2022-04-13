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

func ReSlice[T any](es []interface{}) (ts []T) {
	for _, e := range es {
		ts = append(ts, e.(T))
	}
	return ts
}
