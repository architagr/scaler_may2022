package maximum_unsorted_subarray

func SubUnsort(A []int) []int {
	n := len(A)
	i, j := 0, n-1
	//        find the index from left to right where sorting fail
	for i < n-1 && A[i] <= A[i+1] {
		i++
	}
	// if the i is equal to the array size hence array is already sorted, no need to check subArray size
	if i == j {
		return []int{-1}
	}
	//        find the index from right to left where sorting fail
	for j > 0 && A[j] >= A[j-1] {
		j--
	}

	//      find the minimum and maximum elements in the unsorted subArray [i,j] from i to j
	max, min := A[i+1], A[i+1]
	for k := i; k <= j; k++ {
		max = maxValue(A[k], max)
		min = minValue(A[k], min)
	}
	left, right := 0, n-1
	//      from 0 to i, find the first index where the element is greater than to the minimum
	for left <= i && A[left] <= min {
		left++
	}
	//      from n-1 to j, find the first index where the element is less than  to the maximum
	for right >= j && A[right] >= max {
		right--
	}
	//      return the index
	return []int{left, right}
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
