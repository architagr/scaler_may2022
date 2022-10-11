package perfect_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected string
}

func TestPerfectNumber(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    3,
		expected: "1111",
	})
	testCases = append(testCases, TestCase{
		input:    4,
		expected: "1221",
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := PerfectNumber(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
