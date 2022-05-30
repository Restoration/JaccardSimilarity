package jaccard_similarity

import "sort"

func Intersection(a, b []int) []int {
	s := make(map[int]struct{}, len(a))
	for _, data := range a {
		s[data] = struct{}{}
	}
	r := make([]int, 0, len(a))
	for _, data := range b {
		if _, ok := s[data]; !ok {
			continue
		}
		r = append(r, data)
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}
