package good_subarray

func FindCountOfGoodArray(A []int, B int) int {
	count := 0
	n := len(A)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += A[j]
			if (j-i+1)%2 == 0 {
				if sum < B {
					count++
				}
			} else {
				if sum > B {
					count++
				}
			}
		}
	}
	return count
}
