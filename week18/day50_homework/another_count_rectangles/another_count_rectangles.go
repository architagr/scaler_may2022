package another_count_rectangles

func CountRectangles(A []int, B int) int {
	count, i, j := 0, 0, len(A)-1
	MOD := 1000000007
	for i < j {
		if A[i]*A[j] < B {
			count = (count + (2*(j-i))%MOD) % MOD

			i++
		} else {
			j--
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i]*A[i] < B {
			count = (count + 1) % MOD
		}
	}
	return count
}
