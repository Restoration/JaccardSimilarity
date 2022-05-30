package jaccard_similarity

import (
	"sort"
)

func Union(a, b []int) []int {
	s := make(map[int]struct{}, len(a))
	for _, data := range a {
		s[data] = struct{}{}
	}
	for _, data := range b {
		s[data] = struct{}{}
	}
	r := make([]int, 0, len(s))

	for key := range s {
		r = append(r, key)
	}

	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}
