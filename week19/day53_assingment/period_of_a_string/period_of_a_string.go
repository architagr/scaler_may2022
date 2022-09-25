package period_of_a_string

func PeriodOfString(A string) int {
	n := len(A)
	lps := make([]int, n)
	lps[0] = 0
	for i := 1; i < n; i++ {
		x := lps[i-1]
		for A[x] != A[i] {
			if x == 0 {
				x = -1
				break
			}
			x = lps[x-1]
		}
		lps[i] = x + 1
	}
	return n - lps[n-1]
}
