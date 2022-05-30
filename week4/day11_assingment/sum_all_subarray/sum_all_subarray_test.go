package sum_all_subarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int64
}

func TestSubarraySum(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3},
		expected: 20,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 1, 3},
		expected: 19,
	})

	for _, val := range testCases {
		t.Run(fmt.Sprintf("testing %+v", val.input), func(tb *testing.T) {
			got := SubarraySum(val.input)
			if got != val.expected {
				tb.Errorf("Testing Failed for %+v expected %d but got %d", val.input, val.expected, got)
			}
		})
	}
}
