package subarray_or

func CountSubArrayOr(A []int) int {
	n := len(A) * (len(A) + 1) / 2
	x := 1

	count := 0
	for i := 0; i < 31; i++ {
		if i > 0 {
			x *= 2
		}
		zeroSubArrayCount := countZeroBitsSubArray(A, i)
		count += (x * (n - zeroSubArrayCount)) % 1000000007
	}
	return count % 1000000007
}

func countZeroBitsSubArray(A []int, index int) int {
	x := 1 << index

	count, sum := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i]&x == x {
			sum += count * (count + 1) / 2
			count = 0

		} else {
			count++
		}
	}
	if count > 0 {
		sum += count * (count + 1) / 2
	}
	return sum
}
