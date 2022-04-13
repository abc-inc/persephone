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

package event

import "reflect"

type FormatEvent struct {
	Format string
	Sep    string
}

type Subscriber[E any] func(e E)

var subByType = make(map[string][]Subscriber[any])

func Subscribe[E any](e E, s Subscriber[E]) {
	subByType[typeOf(e)] = append(subByType[typeOf(e)], func(e interface{}) {
		s(e.(E))
	})
}

func Publish[E any](e E) {
	for _, s := range subByType[typeOf(e)] {
		s(e)
	}
}

func typeOf(e interface{}) string {
	return reflect.TypeOf(e).Name()
}
