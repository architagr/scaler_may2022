package pairs_with_xor

func PairsWithXor(A []int, B int) int {
	count, n := 0, len(A)

	mapA := make(map[int]int)

	for i := 0; i < n; i++ {
		if _, ok := mapA[A[i]]; ok {
			count++
		}
		x := A[i] ^ B
		mapA[x] = 1
	}

	return count
}
