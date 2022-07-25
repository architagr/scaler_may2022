package rain_water_trapped

func FindRainWaterTrapped(A []int) int {
	n := len(A)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	max := A[0]
	for i := 1; i < n; i++ {
		if max < A[i] {
			max = A[i]
		}
		leftMax[i] = max
	}
	max = A[n-1]
	for i := n - 2; i >= 0; i-- {
		if max < A[i] {
			max = A[i]
		}
		rightMax[i] = max
	}
	sum := 0
	for i := 0; i < n; i++ {
		water := minValue(leftMax[i], rightMax[i]) - A[i]
		if water > 0 {
			sum += water
		}
	}
	return sum
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
