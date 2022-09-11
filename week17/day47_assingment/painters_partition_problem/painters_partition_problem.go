package painters_partition_problem

func Paint(A, B int, C []int) int {
	l, r := GetArrMaxAndSum(C)
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		if check(A, mid, C) {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return (ans * B) % 10000003
}
func GetArrMaxAndSum(A []int) (max int, sum int) {
	max = A[0]
	sum = A[0]
	for i := 1; i < len(A); i++ {
		if max < A[i] {
			max = A[i]
		}
		sum += A[i]
	}
	return
}

func check(A, mid int, C []int) bool {
	count := 1
	last := 0

	for i := 0; i < len(C); i++ {
		if last+C[i] <= mid {
			last += C[i]
		} else {
			count++
			last = C[i]
		}
	}
	return count <= A
}
