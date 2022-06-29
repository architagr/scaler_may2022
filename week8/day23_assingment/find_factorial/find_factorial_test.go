package find_factorial

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestFindFactorial(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    4,
		expected: 24,
	})
	testCases = append(testCases, TestCase{
		input:    5,
		expected: 120,
	})

	testCases = append(testCases, TestCase{
		input:    0,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    1,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := findFactorial(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
