package window_string

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type TestCase struct {
		inputA, inputB string
		expected       string
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   "ADOBECODEBANC",
		inputB:   "ABC",
		expected: "BANC",
	})

	testCases = append(testCases, TestCase{
		inputA:   "Aa91b",
		inputB:   "ab",
		expected: "a91b",
	})

	testCases = append(testCases, TestCase{
		inputA:   "Aa91b",
		inputB:   "ac",
		expected: "",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MinWindowString(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %s but we got %s", (i + 1), testCase.expected, got)
			}
		})
	}
}
