package subarray_sum_zero

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindSubArrayWithSum0(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, -2, 4, -4},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 2, 4, 4},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, -1},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindSubArrayWithSum0(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
