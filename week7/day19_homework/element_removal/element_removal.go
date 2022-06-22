package element_removal

import (
	"sort"
)

func FindCostOfRemoval(A []int) int {
	n, cost := len(A), 0

	sort.Ints(A)
	for i := 0; i < n; i++ {
		cost += (n - i) * A[i]
	}

	return cost
}
