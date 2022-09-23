package longest_substring_without_repeat

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type TestCase struct {
		input    string
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "aaaaaaaaaabaaaaaaaaa",
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    "abcabcbb",
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    "AaaA",
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := LengthOfLongestSubstring(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
