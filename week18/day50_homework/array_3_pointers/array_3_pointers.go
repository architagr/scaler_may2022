package array_3_pointers

import "math"

func Minimize(A []int, B []int, C []int) int {
	i, j, k := 0, 0, 0
	n, m, o := len(A), len(B), len(C)
	ans := math.MaxInt
	for i < n && j < m && k < o {
		a, b, c := A[i], B[j], C[k]
		max := getMax(a, b, c)
		ans = minValue(ans, max)

		if a <= b && a <= c {
			i++
		} else if b <= a && b <= c {
			j++
		} else if c <= a && c <= b {
			k++
		}
	}
	return ans
}
func getMax(a, b, c int) int {
	ab := int(math.Abs(float64(a - b)))
	bc := int(math.Abs(float64(b - c)))
	ca := int(math.Abs(float64(c - a)))

	return maxValue(maxValue(ab, bc), ca)
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
