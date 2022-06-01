package max_sum_contiguous_subarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestGetMaxSumContiguousSubArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, -10},
		expected: 10,
	})
	testCases = append(testCases, TestCase{
		input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		input:    []int{-500},
		expected: -500,
	})
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.input), func(tb *testing.T) {
			got := GetMaxSumContiguousSubArray(testCase.input)
			if got != testCase.expected {
				tb.Errorf("testing %+v, expected %d, got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
