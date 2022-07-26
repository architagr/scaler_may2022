package add_one_to_number

func PlusOne(A []int) []int {
	for len(A) > 0 && A[0] == 0 {
		A = A[1:]
	}
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] != 9 {
			A[i] += 1
			return A
		}
		A[i] = 0
	}
	return append([]int{1}, A...)
}
