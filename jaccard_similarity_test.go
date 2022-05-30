package jaccard_similarity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJaccardSimilarity(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := make(map[int][]int)
	b[1] = []int{1, 2, 3, 4}
	b[2] = []int{4, 5, 6, 8}
	b[3] = []int{1, 2, 3, 4}
	b[4] = []int{7, 8, 9, 1}
	b[5] = []int{7, 8, 9, 5}
	b[6] = []int{7, 8, 9, 1}

	expected := []*Result{
		{
			Key:   1,
			Point: 0.8,
		},
		{
			Key:   2,
			Point: 0.2857142857142857,
		},
		{
			Key:   3,
			Point: 0.8,
		},
		{
			Key:   4,
			Point: 0.125,
		},
		{
			Key:   5,
			Point: 0.125,
		},
		{
			Key:   6,
			Point: 0.125,
		},
	}

	result := JaccardSimilarity(a, b)
	assert.True(t, reflect.DeepEqual(expected, result))
}
