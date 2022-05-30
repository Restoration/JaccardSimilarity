package jaccard_similarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	expected := []int{3, 4, 5}
	result := Intersection(a, b)
	assert.Equal(t, expected, result)
}
