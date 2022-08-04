package min_XOR_value

import "sort"

func FindMinXor(A []int) int {
	sort.Ints(A)
	min := A[0] ^ A[1]
	for i := 2; i < len(A); i++ {
		x := A[i-1] ^ A[i]
		if x < min {
			min = x
		}
	}
	return min
}
