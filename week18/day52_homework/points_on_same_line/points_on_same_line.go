package points_on_same_line

import "fmt"

func PointsOnSameLine(A, B []int) int {
	hashA := make(map[string]int)

	res := 0

	for i := 0; i < len(A); i++ {
		equal := 1
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] && B[j] == B[i] {
				equal++
			} else {
				y := B[i] - B[j]
				x := A[i] - A[j]
				g__ := gcd(y, x)
				str := fmt.Sprintf("%d_%d", (x / g__), (y / g__))
				hashA[str]++
			}
		}
		res = maxValue(res, equal)
		for _, count := range hashA {
			res = maxValue(res, count+equal)
		}
		hashA = make(map[string]int)
	}
	return res
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
