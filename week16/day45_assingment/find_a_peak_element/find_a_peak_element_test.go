package find_a_peak_element

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindPeakElement(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 4, 5},
		expected: 5,
	})

	testCases = append(testCases, TestCase{
		input:    []int{5, 17, 100, 11},
		expected: 100,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1000000000, 1000000000},
		expected: 1000000000,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindPeakElement(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
