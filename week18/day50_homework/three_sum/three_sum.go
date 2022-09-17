package three_sum

import (
	"math"
	"sort"
)

func ThreeSumClosest(A []int, B int) int {
	n := len(A)
	sort.Ints(A)
	sum := int64(math.MaxInt64)
	for i := 0; i <= n-3; i++ {
		j, k := i+1, n-1
		for j < k {
			s := int64(A[i] + A[j] + A[k])
			if s == int64(B) {
				return B
			}
			if math.Abs(float64(s-int64(B))) < math.Abs(float64(sum-int64(B))) {
				sum = s
			}
			if s < int64(B) {
				j++
			} else {
				k--
			}
		}
	}
	return int(sum)
}
