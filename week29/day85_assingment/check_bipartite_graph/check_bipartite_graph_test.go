package check_bipartite_graph

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   int
	inputB   [][]int
	expected int
}

func TestCycleUndirectedGraph(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   [][]int{{0, 1}},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputA:   3,
		inputB:   [][]int{{0, 1}, {0, 2}, {1, 2}},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
