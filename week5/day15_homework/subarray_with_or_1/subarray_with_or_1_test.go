package subarray_with_or_1

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int64
	inputB   []int
}

func TestFindSubArray(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    3,
		inputB:   []int{1, 0, 1},
		expected: 5,
	})

	testcases = append(testcases, TestCase{
		input:    2,
		inputB:   []int{1, 0},
		expected: 2,
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := FindSubArray(testCase.input, testCase.inputB)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
