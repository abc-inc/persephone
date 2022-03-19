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
