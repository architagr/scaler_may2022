package gray_code

func GrayCode(n int) []int {
	n = 1 << n
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i ^ (i >> 1)
	}
	return arr
}
