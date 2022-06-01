package count_subarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	inputMax int
	expected int
}

func TestFindCountSubarraySumLessthenNumber(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 5, 6},
		inputMax: 10,
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 11, 2, 3, 15},
		inputMax: 10,
		expected: 4,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v for less then %d", testCase.input, testCase.inputMax), func(tb *testing.T) {
			got := FindCountSubarraySumLessthenNumber(testCase.input, testCase.inputMax)

			if got != testCase.expected {
				tb.Errorf("testing %+v for less then %d, expected %d, got %d", testCase.input, testCase.inputMax, testCase.expected, got)
			}
		})
	}
}
