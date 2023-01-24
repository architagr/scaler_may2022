package combination_sum_ii

import (
	"fmt"
	"sort"
)

func combinationSum(A []int, B int) [][]int {
	hashMap := make(map[string]bool)
	result := make([][]int, 0)
	sort.Ints(A)
	check(A, []int{}, B, 0, 0, &result, &hashMap)
	return result
}

func check(A, currentArr []int, sum, currentSum, currentIndex int, result *[][]int, hashMap *map[string]bool) {

	if currentSum == sum {
		str := fmt.Sprintf("%+v", currentArr)
		if _, ok := (*hashMap)[str]; !ok {
			*result = append(*result, append([]int{}, currentArr...))
			(*hashMap)[str] = true
		}
		return
	}
	if currentIndex == len(A) || currentSum > sum {
		return
	}
	currentArr = append(currentArr, A[currentIndex])

	check(A, currentArr, sum, currentSum+A[currentIndex], currentIndex+1, result, hashMap)
	currentArr = currentArr[:len(currentArr)-1]
	check(A, currentArr, sum, currentSum, currentIndex+1, result, hashMap)
}
