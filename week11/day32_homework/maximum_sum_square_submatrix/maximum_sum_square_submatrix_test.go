package maximum_sum_square_submatrix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	b        int
	expected int
}

func TestMaxSqSubMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input: [][]int{{1, 1, 1, 1, 1},
			{2, 2, 2, 2, 2},
			{3, 8, 6, 7, 3},
			{4, 4, 4, 4, 4},
			{5, 5, 5, 5, 5}},
		b:        3,
		expected: 48,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := MaxSqSubMatrix(testCase.input, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
