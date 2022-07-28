package odd_even_subsequences

func FindEvenOddSubSequence(A []int) int {
	isEven := A[0]%2 == 0
	count := 1

	for i := 1; i < len(A); i++ {
		x := A[i]%2 == 0
		if x != isEven {
			isEven = x
			count++
		}
	}
	return count
}
