package specialindex

func solve(A []int) int {
	evenPrefixSum, oddPrefixSum := getEvenPrefixSum(A), getOddPrefixSum(A)
	count := 0
	for i := range A {
		if getEvenSum(evenPrefixSum, oddPrefixSum, i) == getOddSum(evenPrefixSum, oddPrefixSum, i) {
			count++
		}
	}
	return count
}

func getEvenSum(evenPrefixSum, oddPrefixSum []int, index int) int {
	result, n := 0, len(evenPrefixSum)-1
	if index > 0 {
		result = evenPrefixSum[index-1]
	}
	result += oddPrefixSum[n] - oddPrefixSum[index]

	return result
}
func getOddSum(evenPrefixSum, oddPrefixSum []int, index int) int {
	result, n := 0, len(evenPrefixSum)-1
	if index > 0 {
		result = oddPrefixSum[index-1]
	}
	result += evenPrefixSum[n] - evenPrefixSum[index]

	return result
}

func getEvenPrefixSum(input []int) []int {
	result := make([]int, len(input))
	result[0] = input[0]
	for i := 1; i < len(input); i++ {
		result[i] = result[i-1]
		if i&1 != 1 {
			result[i] += input[i]
		}
	}
	return result
}

func getOddPrefixSum(input []int) []int {
	result := make([]int, len(input))
	result[0] = 0
	for i := 1; i < len(input); i++ {
		result[i] = result[i-1]
		if i&1 == 1 {
			result[i] += input[i]
		}
	}
	return result
}
