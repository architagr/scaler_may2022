package closest_palindrome

import (
	"fmt"
	"testing"
)

func TestSum2II(t *testing.T) {
	type TestCase struct {
		input, expected string
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "aaaaaaaaaabaaaaaaaaa",
		expected: "YES",
	})

	testCases = append(testCases, TestCase{
		input:    "asdfghjklasdfghjklasdfgqjklasdfghjkllkjhgfdsalkjhgfdsalkjhgfdsalkjhgfdsa",
		expected: "YES",
	})

	testCases = append(testCases, TestCase{
		input:    "abbba",
		expected: "YES",
	})

	testCases = append(testCases, TestCase{
		input:    "adaddb",
		expected: "NO",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CheckClosestPalindrome(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %s but we got %s", (i + 1), testCase.expected, got)
			}
		})
	}
}
