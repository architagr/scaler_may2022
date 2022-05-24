package goodpair

func CheckGoodPairExists(A []int, B int) int {
	for i := 1; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i]+A[j] == B {
				return 1
			}
		}
	}
	return 0
}
