package combination_sum

import "sort"

func combinationSum(A []int, B int) [][]int {
	A = removeDuplicate(A)
	result := make([][]int, 0)
	sort.Ints(A)
	check(A, []int{}, B, 0, 0, &result)
	return result
}
func removeDuplicate(arr []int) []int {
	hashmap := make(map[int]bool)
	result := make([]int, 0, len(arr))
	for _, val := range arr {
		hashmap[val] = true
	}
	for key := range hashmap {
		result = append(result, key)
	}
	return result
}
func check(A, currentArr []int, sum, currentSum, currentIndex int, result *[][]int) {
	if currentIndex == len(A) || currentSum > sum {
		return
	}
	if currentSum == sum {
		*result = append(*result, append([]int{}, currentArr...))
		return
	}
	currentArr = append(currentArr, A[currentIndex])

	check(A, currentArr, sum, currentSum+A[currentIndex], currentIndex, result)
	currentArr = currentArr[:len(currentArr)-1]
	check(A, currentArr, sum, currentSum, currentIndex+1, result)
}
