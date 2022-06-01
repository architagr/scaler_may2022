package count_subarray

func FindCountSubarraySumLessthenNumber(A []int, B int) int {
	count := 0
	sum := 0

	for i := 0; i < len(A); i++ {
		sum = 0

		for j := i; j < len(A); j++ {
			sum += A[j]

			if sum < B {
				count++
			}
		}
	}
	return count
}
