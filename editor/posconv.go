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

package editor

// PosConv provides line and column info for absolute positions in a multi-line
// string and vice versa.
type PosConv struct {
	newLines []int
}

// NewPosConv splits the given input line by line and indexes their position.
func NewPosConv(input string) *PosConv {
	pc := &PosConv{}
	for i, s := range input {
		if s == '\n' {
			pc.newLines = append(pc.newLines, i)
		}
	}
	return pc
}

// ToAbsolute calculates the absolute position of line and column.
func (pc PosConv) ToAbsolute(line, column int) int {
	if line < 2 {
		return column
	}
	return pc.newLines[line-2] + column + 1
}

// ToRelative determines line and column for a given position.
func (pc PosConv) ToRelative(abs int) (int, int) {
	for i := len(pc.newLines) - 1; i >= 0; i-- {
		column := abs - pc.newLines[i]
		if column >= 1 {
			return i + 2, column - 1
		}
	}
	return 1, abs
}
