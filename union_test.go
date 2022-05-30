package jaccard_similarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	result := Union(a, b)
	assert.Equal(t, expected, result)
}
