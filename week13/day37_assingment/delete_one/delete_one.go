package delete_one

import "scaler-may-22/week13/day37_assingment/greatest_common_divisor"

func DeleteOne(A []int) int {
	prefix := prefixGCD(A)
	surfix := surfixGCD(A)
	n := len(A)
	max := surfix[1]

	for i := 1; i < n; i++ {
		if i == n-1 {
			max = maxValue(max, prefix[i-1])
		} else {
			max = maxValue(max, greatest_common_divisor.Gcd(prefix[i-1], surfix[i+1]))
		}
	}
	return max
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func prefixGCD(A []int) []int {
	result := make([]int, len(A))

	result[0] = A[0]
	for i := 1; i < len(A); i++ {
		result[i] = greatest_common_divisor.Gcd(result[i-1], A[i])
	}
	return result
}

func surfixGCD(A []int) []int {
	result := make([]int, len(A))
	n := len(A)
	result[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		result[i] = greatest_common_divisor.Gcd(result[i+1], A[i])
	}
	return result
}
