package jaccard_similarity

import (
	"sort"
)

func jaccard(a []int, b []int) float64 {
	intersectionResult := Intersection(a, b)
	countIntersection := len(intersectionResult)

	unionResult := Union(a, b)
	countUnion := len(unionResult)

	return float64(countIntersection) / float64(countUnion)
}

func JaccardSimilarity(t []int, c map[int][]int) []*Result {
	var r []*Result
	for i, v := range c {
		p := jaccard(t, v)
		r = append(r, &Result{Key: i, Point: p})
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Key < r[j].Key
	})
	return r
}
