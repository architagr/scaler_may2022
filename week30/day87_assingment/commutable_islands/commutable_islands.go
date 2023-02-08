package commutable_islands

import (
	"sort"
)

func ComutableIslands(A int, B [][]int) int {
	sort.Slice(B, func(i, j int) bool {
		return sortNode(B[i], B[j])
	})
	ans := 0
	componentArray := make([]int, A+1)
	for i := 0; i < len(componentArray); i++ {
		componentArray[i] = i
	}
	for i := 0; i < len(B); i++ {
		pa, pb := parentNode(B[i][0], componentArray), parentNode(B[i][1], componentArray)
		if pa != pb {
			ans += B[i][2]
			componentArray[maxValue(pa, pb)] = minValue(pa, pb)
		}
	}
	return ans
}
func maxValue(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func sortNode(a, b []int) bool {
	return a[2] < b[2]
}

func parentNode(node int, componentArr []int) int {
	if componentArr[node] == node {
		return node
	}
	return parentNode(componentArr[node], componentArr)
}
