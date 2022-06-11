package subarray_with_or_1

func FindSubArray(A int, B []int) int64 {
	var result, zeroCount int64 = 0, 0

	for start := 0; start < A; start++ {
		if B[start] == 1 {
			result += int64(A - start)
			result += zeroCount * int64(A-start)
			zeroCount = 0
		} else {
			zeroCount++
		}
	}

	return result
}
