package cyclic_permutations

func CyclicPermutation(A string, B string) int {
	n := len(A)
	count := 0
	for i := 0; i < n; i++ {
		B = string(B[1:]) + string(B[0])
		if B == A {
			count++
		}
	}
	return count
}
