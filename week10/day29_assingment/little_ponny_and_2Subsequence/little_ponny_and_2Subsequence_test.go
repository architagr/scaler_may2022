package little_ponny_and_2Subsequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected string
}

func TestFindLexicographicalMinString(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "abcdsfhjagj",
		expected: "aa",
	})
	testCases = append(testCases, TestCase{
		input:    "ksdjgha",
		expected: "da",
	})
	testCases = append(testCases, TestCase{
		input:    "djjhibvetj",
		expected: "be",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindLexicographicalMinString(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %s but got %s", (i + 1), testCase.expected, got)
			}
		})
	}
}
