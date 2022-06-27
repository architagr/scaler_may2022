package palindrome

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestCheckPalindrome(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "abcde",
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		input:    "abbaee",
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    "abbadee",
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CheckPalindrome(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
