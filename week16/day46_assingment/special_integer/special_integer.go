package special_integer

func FindSpecialInt(A []int, B int) int {
	l, ans, sum := 0, len(A), 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		for sum > B {
			sum -= A[l]
			l++
			ans = minValue(ans, i-l+1)
		}
	}
	return ans
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
