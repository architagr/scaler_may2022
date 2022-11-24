package min_sum_path_in_matrix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestMinPathSum(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]int{
			{20, 29, 84, 4, 32, 60, 86, 8, 7, 37},
			{77, 69, 85, 83, 81, 78, 22, 45, 43, 63},
			{60, 21, 0, 94, 59, 88, 9, 54, 30, 80},
			{40, 78, 52, 58, 26, 84, 47, 0, 24, 60},
			{40, 17, 69, 5, 38, 5, 75, 59, 35, 26},
			{64, 41, 85, 22, 44, 25, 3, 63, 33, 13},
			{2, 21, 39, 51, 75, 70, 76, 57, 56, 22},
			{31, 45, 47, 100, 65, 10, 94, 96, 81, 14},
		},
		expected: 470,
	})
	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 3, 2},
			{4, 3, 1},
			{5, 6, 1},
		},
		expected: 8,
	})

	// testCases = append(testCases, TestCase{
	// 	input: [][]int{
	// 		{1, -3, 2},
	// 		{2, 5, 10},
	// 		{5, -5, 1},
	// 	},
	// 	expected: -1,
	// })

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MinPathSum(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
