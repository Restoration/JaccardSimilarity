# JaccardSimilarity

JaccardSimilarity is a collaborative filtering logic package.  
You can get the similarity index.

## Example
```go
package main

import (
	"fmt"

	"github.com/Restoration/jaccard_similarity"
)

func main() {
	target := []int{1, 2, 3, 4, 5}
	comparison := make(map[int][]int)
	comparison[1] = []int{1, 2, 3, 4}
	comparison[2] = []int{4, 5, 6, 8}
	comparison[3] = []int{1, 2, 3, 4}
	comparison[4] = []int{7, 8, 9, 1}
	comparison[5] = []int{7, 8, 9, 5}
	comparison[6] = []int{7, 8, 9, 1}
	result := JaccardSimilarity(target, comparison)
	for _, v := range result {
		fmt.Printf("%v's similarity index is %v\n", v.Key, v.Point)
	}
}
```

### Result
```
1's similarity index is 0.8
2's similarity index is 0.2857142857142857
3's similarity index is 0.8
4's similarity index is 0.125
5's similarity index is 0.125
6's similarity index is 0.125
```

## License
BSD 2-Clause