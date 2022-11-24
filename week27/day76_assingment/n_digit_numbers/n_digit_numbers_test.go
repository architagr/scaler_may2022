package n_digit_numbers

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB int
	expected       int
}

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   4,
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		inputA:   1,
		inputB:   3,
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputA:   56,
		inputB:   17,
		expected: 4692503,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
