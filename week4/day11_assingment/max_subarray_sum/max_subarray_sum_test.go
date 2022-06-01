package max_subarray_sum

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	inputMax int
	expected int
}

func TestFindMaxSubArraySumBelowValue(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 1, 3, 4, 5},
		inputMax: 12,
		expected: 12,
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, 2, 2},
		inputMax: 1,
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		input:    []int{4},
		inputMax: 75,
		expected: 4,
	})
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v for less then %d", testCase.input, testCase.inputMax), func(tb *testing.T) {
			got := FindMaxSubArraySumBelowValue(len(testCase.input), testCase.inputMax, testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %+v for less then %d, expected %d, got %d", testCase.input, testCase.inputMax, testCase.expected, got)
			}
		})
	}
}
