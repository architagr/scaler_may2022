package longest_consecutive_sequence

import (
	"fmt"
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{100, 4, 200, 1, 3, 2},
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 3, 2},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 3},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := LongestConsecutiveSequence(testCase.input)
			if got != testCase.expected {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
