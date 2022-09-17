package shaggy_and_distances

import (
	"fmt"
	"testing"
)

func TestMinDistanceSpecialPair(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{7, 1, 3, 4, 1, 7},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MinDistanceSpecialPair(testCase.input)
			if got != testCase.expected {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
