// Copyright 2022 The Persephone authors
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

package lang

import (
	"regexp"
	"strings"
)

// EscapeCypher uses backticks "`" to escape a string, which contains special
// characters.
func EscapeCypher(str string) string {
	prefix := ""
	if strings.HasPrefix(str, ":") {
		prefix = ":"
	}

	content := strings.TrimPrefix(str, prefix)
	if ok, err := regexp.MatchString(`^[A-Za-z][A-Za-z0-9_]*$`, content); err == nil && ok {
		return str
	}
	return prefix + "`" + strings.ReplaceAll(content, "`", "``") + "`"
}
