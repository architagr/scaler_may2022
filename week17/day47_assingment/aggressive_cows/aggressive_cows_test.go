package aggressive_cows

import (
	"fmt"
	"testing"
)

func TestAggressiveCows(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, 5},
		b:        3,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    []int{5, 17, 100, 11},
		b:        2,
		expected: 95,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := AggressiveCows(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
