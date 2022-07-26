package max_sum_contiguous_subarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{-163, -20},
		expected: -20,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, -10},
		expected: 10,
	})
	testCases = append(testCases, TestCase{
		input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := MaxSubArray(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %d, got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
