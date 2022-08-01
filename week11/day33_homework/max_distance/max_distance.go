package max_distance

func MaxDistance(A []int) int {
	n := len(A)
	max := 0
	rightMax := make([]int, n)
	rightMax[n-1] = A[n-1]
	for j := n - 2; j >= 0; j-- {
		rightMax[j] = maxValue(rightMax[j+1], A[j])
	}

	for i, j := 0, 0; i < n && j < n; {
		if rightMax[j] >= A[i] {
			max = maxValue(max, j-i)
			j++
		} else {
			i++
		}
	}
	return max
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
