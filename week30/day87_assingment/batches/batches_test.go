package batches

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputD int
	inputB         []int
	inputC         [][]int
	expected       int
}

func TestSolve(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   7,
		inputD:   12,
		inputB:   []int{1, 6, 7, 2, 9, 4, 5},
		inputC:   [][]int{{1, 2}, {2, 3}, {5, 6}, {5, 7}},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		inputA:   5,
		inputD:   6,
		inputB:   []int{1, 2, 3, 4, 5},
		inputC:   [][]int{{1, 5}, {2, 3}},
		expected: 1,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := countBatches(testCase.inputA, testCase.inputB, testCase.inputC, testCase.inputD)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
