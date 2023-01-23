package number_of_squareful_arrays

import "math"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */

var hashMap map[int]bool

func solve(A []int) int {
	if len(A) == 1 {
		return 0
	}

	hashMap = make(map[int]bool)
	generate(A, 0)
	return len(hashMap)
}

func generate(A []int, start int) {
	if start == len(A) {
		hashMap[createNum(A)] = true
		return
	}
	for i := start; i < len(A); i++ {
		A[i], A[start] = A[start], A[i]
		if start == 0 || (start > 0 && checkPerfectSq(A[start]+A[start-1])) {
			generate(A, start+1)
		}
		A[i], A[start] = A[start], A[i]
	}
}
func checkPerfectSq(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}
func createNum(A []int) int {
	num := 0
	for _, val := range A {
		num = (num * 10) + val
	}
	return num
}
