package knapsack_ii_0_1

/**
 * @input A : Integer array
 * @input B : Integer array
 * @input C : Integer
 *
 * @Output Integer
 */

func Solve(A []int, B []int, C int) int {
	return knapsackBottomUpOptimizeSpace(A, B, len(A), C)
}

func knapsackBottomUpOptimizeSpace(values, weights []int, n, capacity int) int {
	prev := make([]int, capacity+1)
	curr := make([]int, capacity+1)
	for i := 0; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if i == 0 || w == 0 {
				curr[w] = 0
			} else if weights[i-1] > w {
				curr[w] = prev[w]
			} else {
				curr[w] = maxValue(values[i-1]+prev[w-weights[i-1]], prev[w])
			}
		}
		prev = curr
		curr = make([]int, capacity+1)
	}
	return prev[capacity]
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
