package find_fibonacci_recursion

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestFindAthFibonacci(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    2,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    9,
		expected: 34,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := findAthFibonacci(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
