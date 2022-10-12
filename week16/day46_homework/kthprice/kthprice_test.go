package kthprice

import (
	"fmt"
	"testing"
)

func TestKthPrice(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{2, 1, 4, 3, 2},
		b:        3,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    []int{74, 90, 85, 58, 69, 77, 90, 85, 18, 36},
		b:        1,
		expected: 18,
	})

	testCases = append(testCases, TestCase{
		input:    []int{60, 94, 63, 3, 86, 40, 93, 36, 56, 48, 17, 10, 23, 43, 77, 1, 1, 93, 79, 4, 10, 47, 1, 99, 91, 53, 99, 18, 52, 61, 84, 10, 13, 52, 3, 9, 78, 16, 7, 62},
		b:        22,
		expected: 52,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := KthPrice(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
