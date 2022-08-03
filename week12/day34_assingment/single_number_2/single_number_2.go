package single_number_2

func FindSingleNumber(A []int) int {
	one, two := 0, 0

	for i := 0; i < len(A); i++ {
		x := A[i]

		one = (one ^ x) & (^two)
		two = (two ^ x) & (^one)
	}

	return one
}
