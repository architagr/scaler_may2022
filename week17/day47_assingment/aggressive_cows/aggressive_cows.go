package aggressive_cows

import "sort"

func AggressiveCows(A []int, B int) int {
	n := len(A)
	sort.Ints(A)
	l, r := 1, A[n-1]-A[0]
	ans := 0

	for l <= r {
		mid := (l + r) / 2
		if Check(A, mid, B) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
func Check(A []int, size, B int) bool {
	last := A[0]
	count := 1

	for i := 1; i < len(A); i++ {
		if last+size <= A[i] {
			last = A[i]
			count++
			if count == B {
				return true
			}
		}
	}
	return false
}
