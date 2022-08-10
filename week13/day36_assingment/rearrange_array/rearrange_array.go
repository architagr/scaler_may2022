package rearrange_array

func RearrangeArray(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		A[i] *= n
	}
	for i := 0; i < n; i++ {
		A[i] += A[A[i]/n] / n
	}

	for i := 0; i < n; i++ {
		A[i] %= n
	}

}
