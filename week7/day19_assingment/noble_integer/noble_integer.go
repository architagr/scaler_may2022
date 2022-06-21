package noble_integer

import (
	"sort"
)

func CheckIfNobleExist(A []int) int {
	sort.Ints(A)
	n := len(A)
	if A[n-1] == 0 {
		return 1
	}
	for i := n - 2; i >= 0; i-- {
		if A[i] != A[i+1] && A[i] == (n-i-1) {
			return 1
		}
	}
	return -1
}
