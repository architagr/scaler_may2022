package path_in_directed_graph

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   int
	inputB   [][]int
	expected int
}

func TestSolve(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   5,
		inputB:   [][]int{{1, 2}, {4, 1}, {2, 4}, {3, 4}, {5, 2}, {1, 3}},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		inputA:   5,
		inputB:   [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := PathDirectedGraph(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
