package max_sum_without_adjacent_elements

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestNQueens(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 2, 3, 4},
			{2, 3, 4, 5},
		},
		expected: 8,
	})

	testCases = append(testCases, TestCase{
		input: [][]int{
			{1},
			{2},
		},
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input: [][]int{
			{16, 5, 54, 55, 36, 82, 61, 77, 66, 61},
			{31, 30, 36, 70, 9, 37, 1, 11, 68, 14},
		},
		expected: 321,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := adjacent(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
