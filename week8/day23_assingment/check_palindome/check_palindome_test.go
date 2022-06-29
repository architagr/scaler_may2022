package check_palindome

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
		input:    "aba",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "abccba",
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    "abca",
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		input:    "abcd",
		expected: 0,
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
