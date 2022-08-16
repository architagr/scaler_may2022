package prime_sum

import "math"

func PrimeSum(A int) []int {
	nums := make([]bool, A+1)
	sum := make(map[int]int)

	for i := range nums {
		nums[i] = true
	}

	nums[1] = false
	for i := 2; i <= int(math.Sqrt(float64(A))); i++ {
		if nums[i] {
			for j := i * i; j <= A; j = j + i {
				nums[j] = false
			}
		}
	}
	ans := make([]int, 2)
	for i := A; i >= 2; i-- {
		if nums[i] {
			if A-i == i {
				ans[0] = i
				ans[1] = i
			} else {
				if val, ok := sum[i]; ok {
					ans[0] = i
					ans[1] = val
				} else {
					sum[A-i] = i
				}
			}
		}
	}
	if ans[0] == 0 {
		return []int{}
	}
	return ans
}
