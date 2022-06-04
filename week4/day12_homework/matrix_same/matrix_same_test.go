package matrix_same

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB [][]int
	expected       int
}

func TestCompareMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		inputB:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		inputB:   [][]int{{2, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := CompareMatrix(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested case number %d, expected %d and got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
