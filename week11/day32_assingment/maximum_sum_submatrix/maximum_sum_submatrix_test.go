package maximum_sum_submatrix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestMaxSubMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{-6, -6}, {-29, -8}, {3, -8}, {-15, 2}, {25, 25}, {20, -5}},
		expected: 65,
	})

	testCases = append(testCases, TestCase{
		input:    [][]int{{-17, -2}, {20, 10}},
		expected: 30,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := MaxSubMatrix(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
