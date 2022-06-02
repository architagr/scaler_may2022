package reversearray

func ReverseArray(A []int, i, j int) []int {
	for i <= j {
		temp := A[i]
		A[i] = A[j]
		A[j] = temp
		i++
		j--
	}

	return A
}
