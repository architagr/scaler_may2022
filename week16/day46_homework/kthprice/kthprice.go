package kthprice

func KthPrice(A []int, B int) int {
	min, max := A[0], A[0]

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	left, right := min, max
	for left <= right {
		mid := (left + right) / 2
		c, d := verify(A, mid)
		if (c-d+1) <= B && c >= B && d > 0 {
			return mid
		} else if c >= B {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func verify(A []int, mid int) (int, int) {
	count := 0
	duplicateCount := 0
	for _, v := range A {
		if v == mid {
			duplicateCount++
		}
		if v <= mid {
			count++
		}
	}
	return count, duplicateCount
}
