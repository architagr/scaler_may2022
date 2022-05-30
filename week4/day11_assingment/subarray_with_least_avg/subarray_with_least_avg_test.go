package subarray_with_least_avg

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	inputB   int
	expected int
}

func TestSubArrayWithLeastAverage(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{3, 7, 90, 20, 10, 50, 40},
		inputB:   3,
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{3, 7, 5, 20, -10, 0, 12},
		inputB:   2,
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		input:    []int{20, 3, 13, 5, 10, 14, 8, 5, 11, 9, 1, 11},
		inputB:   9,
		expected: 3,
	})

	for _, val := range testCases {
		t.Run(fmt.Sprintf("testing %+v", val.input), func(tb *testing.T) {
			got := SubArrayWithLeastAverage(val.input, val.inputB)
			if got != val.expected {
				tb.Errorf("Testing Failed for %+v expected %d but got %d", val.input, val.expected, got)
			}
		})
	}
}
