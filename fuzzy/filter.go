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
	"sort"
)

type ScoredCandidate[T any] struct {
	item  T
	score float64
}

func filter[T any](items []T, query string, queryHasSlashes bool, keyEx func(T) string) (candidates []T) {
	var scoredCandidates []ScoredCandidate[T]
	for _, i := range items {
		s := keyEx(i)
		if s == "" {
			continue
		}

		score := score(s, query)
		if !queryHasSlashes {
			score = basenameScore(s, query, score)
		}
		if score > 0 {
			scoredCandidates = append(scoredCandidates, ScoredCandidate[T]{i, score})
		}
	}

	sort.Slice(scoredCandidates, func(i, j int) bool {
		return scoredCandidates[i].score > scoredCandidates[j].score
	})
	for _, sc := range scoredCandidates {
		candidates = append(candidates, sc.item)
	}
	return
}
