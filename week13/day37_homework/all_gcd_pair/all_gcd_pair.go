package all_gcd_pair

import "math"

func AllGCDPairs(A []int) []int {
	n := len(A)
	m := int(math.Sqrt(float64(n)))
	ans := make([]int, 0, m)
	max := math.MinInt32
	step := 0
	for i := 0; i < n; i++ {
		max = MaxValue(max, A[i])
		step++
		if step == m {
			ans = append(ans, max)
			max = math.MinInt32
			step = 0
		}
	}
	return ans
}
func MaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
