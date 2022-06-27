package sub_array_given_sum

func SubarrayGivenSum(A []int, B int) []int {
	mapA := make(map[int]int)
	sum, n := 0, len(A)
	for i := 0; i < n; i++ {
		sum += A[i]
		if x, ok := mapA[sum-B]; ok {
			return A[x+1 : i+1]
		} else if sum == B {
			return A[:i+1]
		} else {
			mapA[sum] = i
		}
	}
	return []int{-1}
}
