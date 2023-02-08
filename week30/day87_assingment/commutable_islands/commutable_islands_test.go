package commutable_islands

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   int
	inputB   [][]int
	expected int
}

func TestComutableIslands(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   4,
		inputB:   [][]int{{1, 2, 1}, {2, 3, 4}, {1, 4, 3}, {4, 3, 2}, {1, 3, 10}},
		expected: 6,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := ComutableIslands(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
