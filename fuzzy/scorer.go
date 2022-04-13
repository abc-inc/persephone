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

package fuzzy

import (
	"os"
	"strings"
	"unicode"
)

func basenameScore(str, query string, sc float64) float64 {
	index := len(str) - 1
	for str[index] == os.PathSeparator {
		index--
	}

	var base string
	slashCount := 0
	lastChar := index
	for index >= 0 {
		if str[index] == os.PathSeparator {
			slashCount++
			if base == "" {
				base = str[index+1 : lastChar+1]
			}
		} else if index == 0 {
			if lastChar < len(str)-1 {
				if base == "" {
					base = str[0 : lastChar+1]
				}
			} else {
				if base == "" {
					base = str
				}
			}
		}
		index--
	}

	if base == str {
		sc *= 2
	} else if base != "" {
		sc += score(base, query)
	}
	segmentCount := slashCount + 1
	depth := max(1, 10-segmentCount)
	sc *= float64(depth) * 0.01
	return sc
}

func score(str, query string) float64 {
	if str == query {
		return 1
	}
	if queryIsLastPathSegment(str, query) {
		return 1
	}
	totalCharScore := 0.0
	strLen := len(str)
	indexInString := 0
	for indexInQuery := 0; indexInQuery < len(query); indexInQuery++ {
		c := rune(query[indexInQuery])
		lowerCaseIndex := strings.IndexRune(str, unicode.ToLower(c))
		upperCaseIndex := strings.IndexRune(str, unicode.ToUpper(c))
		minIndex := min(lowerCaseIndex, upperCaseIndex)
		if minIndex == -1 {
			minIndex = max(lowerCaseIndex, upperCaseIndex)
		}
		indexInString = minIndex
		if indexInString == -1 {
			return 0
		}

		charScore := 0.1
		if str[indexInString] == byte(c) {
			charScore += 0.1
		}
		if indexInString == 0 || str[indexInString-1] == os.PathSeparator {
			charScore += 0.8
		} else if c == '-' || c == '_' || c == ' ' {
			charScore += 0.7
		}
		str = str[indexInString+1:]
		totalCharScore += charScore
	}
	queryScore := totalCharScore / float64(len(query))
	return ((queryScore * (float64(len(query)) / float64(strLen))) + queryScore) / float64(2)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
