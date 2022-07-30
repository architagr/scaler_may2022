package beggars_outside_temple

func BeggarsOutsideTemple(A int, B [][]int) []int {
	arr := make([]int, A)

	for i := 0; i < len(B); i++ {
		start := B[i][0] - 1
		end := B[i][1]
		amount := B[i][2]

		arr[start] += amount
		if end < A {
			arr[end] -= amount
		}
	}
	for i := 1; i < A; i++ {
		arr[i] += arr[i-1]
	}
	return arr
}
