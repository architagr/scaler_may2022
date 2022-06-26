package longest_common_prefix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []string
	expected string
}

func TestFindLongestCommonPrefix(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []string{"abcdefgh", "aefghijk", "abcefgh"},
		expected: "a",
	})
	testCases = append(testCases, TestCase{
		input:    []string{"abab", "ab", "abcd"},
		expected: "ab",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindLongestCommonPrefix(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}

}
