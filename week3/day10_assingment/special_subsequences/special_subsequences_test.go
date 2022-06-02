package special_subsequences

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestFindSpecialSubsequences(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "ABCGAG",
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    "GAB",
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %s", testCase.input), func(tb *testing.T) {
			got := FindSpecialSubsequences(testCase.input)
			if got != testCase.expected {
				tb.Errorf("Tested %s extected %d and got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
