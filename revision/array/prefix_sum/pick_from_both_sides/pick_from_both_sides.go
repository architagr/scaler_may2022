package pickfrombothsides

import "math"

func solve(A []int, B int) int {
	result, prefixSum, n := math.MinInt, getPrefixSum(A), len(A)-1
	if B == len(A) {
		return int(prefixSum[n])
	}
	for i := 0; i <= B; i++ {
		sum := 0
		if i > 0 {
			sum = prefixSum[i-1]
		}

		sum += prefixSum[n] - prefixSum[n-B+i]
		if sum > result {
			result = sum
		}
	}

	return result
}

func getPrefixSum(input []int) []int {
	result := make([]int, len(input))
	result[0] = input[0]
	for i := 1; i < len(input); i++ {
		result[i] = result[i-1] + input[i]
	}
	return result
}
