package unique_paths_in_a_grid

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input: [][]int{
			{0, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 0},
		},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := UniquePathsWithObstacles(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
