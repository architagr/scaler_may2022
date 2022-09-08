package food_packets_distribution

func countShops(A []int, num int) int {
	shops_allotted := 0
	for _, population := range A {
		shops_allotted += population / num
	}
	return shops_allotted
}
func sumAndMin(A []int) (int, int) {
	min := A[0]
	sum := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		sum += A[i]
	}
	return min, sum
}
func FindMaxMinCount(A []int, B int) int {
	right, sum := sumAndMin(A)
	if sum < B {
		return 0
	}
	left, ans := 1, 0
	for left <= right {
		mid := left + (right-left)/2
		result := countShops(A, mid)
		if result < B {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}
