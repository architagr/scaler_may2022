package reversearray

func ReverseArray(A []int) []int {
	i := 0
	j := len(A) - 1
	for i <= j {
		temp := A[i]
		A[i] = A[j]
		A[j] = temp
		i++
		j--
	}

	return A
}
