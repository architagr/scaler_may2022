package subarray_or

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected int
}

func TestCountSubArrayOr(t *testing.T) {

	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{7, 8, 9, 10},
		expected: 110,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4, 5},
		expected: 71,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CountSubArrayOr(testCase.inputA)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
