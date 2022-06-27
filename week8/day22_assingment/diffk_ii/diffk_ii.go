package diffk_ii

func Diffk2(A []int, B int) int {
	mapA := make(map[int]int)
	n := len(A)

	for i := 0; i < n; i++ {
		if _, ok := mapA[B+A[i]]; ok {
			return 1
		}
		if _, ok := mapA[A[i]-B]; ok {
			return 1
		}
		if _, ok := mapA[A[i]]; !ok {
			mapA[A[i]] = i
		}
	}
	return 0
}
