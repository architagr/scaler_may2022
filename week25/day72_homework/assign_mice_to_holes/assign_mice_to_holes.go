package assign_mice_to_holes

import (
	"math"
	"sort"
)

func mice(A []int, B []int) int {
	sort.Ints(A)
	sort.Ints(B)
	time := 0
	for i := 0; i < len(A); i++ {
		x := int(math.Abs(float64(A[i] - B[i])))
		time = maxValue(time, x)
	}
	return time
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
