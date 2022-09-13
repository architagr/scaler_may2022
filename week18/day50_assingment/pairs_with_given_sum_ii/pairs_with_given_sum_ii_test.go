package pairs_with_given_sum_ii

import (
	"fmt"
	"testing"
)

func TestSum2II(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 1},
		b:        2,
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 5, 6, 7, 8, 10},
		b:        15,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 5, 5, 7, 8, 10},
		b:        15,
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Sum2II(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
