package knapsack_0_1

var knapsack01 [][]int

func Solve(A []int, B []int, C int) int {
	knapsack01 = make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		knapsack01[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			knapsack01[i][j] = -1
		}
	}
	return knapsack01Recursion(A, B, len(A), C)
}

func knapsack01Recursion(value, weight []int, index, totalWeight int) int {

	if index == 0 || totalWeight == 0 {
		knapsack01[index][totalWeight] = 0
		return 0
	}
	if knapsack01[index][totalWeight] != -1 {
		return knapsack01[index][totalWeight]
	}

	if totalWeight >= weight[index-1] {
		y := value[index-1] + knapsack01Recursion(value, weight, index-1, totalWeight-weight[index-1])
		x := knapsack01Recursion(value, weight, index-1, totalWeight)
		knapsack01[index][totalWeight] = maxValue(x, y)
	} else {
		knapsack01[index][totalWeight] = knapsack01Recursion(value, weight, index-1, totalWeight)
	}
	return knapsack01[index][totalWeight]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
