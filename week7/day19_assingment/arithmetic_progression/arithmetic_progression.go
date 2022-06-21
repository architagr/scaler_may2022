package arithmetic_progression

import (
	"sort"
)

func IfAp(A []int) int {
	sort.Ints(A)
	n := len(A)
	if n > 2 {
		lastDiff := A[1] - A[0]
		for i := 2; i < n; i++ {
			if A[i]-A[i-1] != lastDiff {
				return 0
			}
		}
	}
	return 1
}
