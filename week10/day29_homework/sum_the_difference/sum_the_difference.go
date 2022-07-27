package sum_the_difference

import "sort"

func SumOfDiff(A []int) int {
	sort.Ints(A)
	n, max, min := len(A), 0, 0
	for i := 0; i < n; i++ {
		x := getPow(2, (i), 1000000007)
		y := getPow(2, (n - i - 1), 1000000007)

		max += A[i] % 1000000007 * x
		min += A[i] % 1000000007 * y
	}

	return (max - min + 1000000007) % 1000000007
}
func getPow(A, B, C int) int {
	if C == 1 {
		return 0
	} else if B == 0 {
		return 1
	}
	A = A % C
	if B%2 == 0 {
		return getPow((A*A)%C, B/2, C)
	} else {
		return (A * getPow((A*A)%C, B/2, C)) % C
	}
}
