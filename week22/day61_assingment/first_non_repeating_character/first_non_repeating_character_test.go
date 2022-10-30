package first_non_repeating_character

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestFirstNonRepeatingChar(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "abadbc",
		expected: "aabbdd",
	})
	testCases = append(testCases, TestCase{
		input:    "abcabc",
		expected: "aaabc#",
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := FirstNonRepeatingChar(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
