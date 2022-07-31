package maximum_submatrix_sum_sorted_matrix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int64
}

func TestFindMaxSubmatrixSum(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{-83, -73, -70, -61}, {-56, -48, -13, 4}, {38, 48, 71, 71}},
		expected: 228,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{-5, -4, -3}, {-1, 2, 3}, {2, 2, 4}},
		expected: 12,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindMaxSubmatrixSum(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
