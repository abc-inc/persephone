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

package types

type Data struct {
	Type              Type
	Path              []string
	FilterLastElement bool
}

// AllCompData is the default.
var AllCompData []Data

var AllComp = []Type{Variable, Parameter, PropertyKey, FunctionName, Keyword}

func init() {
	for _, t := range AllComp {
		AllCompData = append(AllCompData, Data{Type: t})
	}
}
