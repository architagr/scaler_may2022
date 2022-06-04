package minor_diagonal_sum

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestMinorDiagonalSum(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}, {1, 5, 9, 8}},
		expected: 14,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := MinorDiagonalSum(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested case number %d, expected %d and got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
