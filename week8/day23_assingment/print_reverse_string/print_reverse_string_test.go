package print_reverse_string

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestReverseString(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "abcd",
		expected: "dcba",
	})

	testCases = append(testCases, TestCase{
		input:    "scaleracademy",
		expected: "ymedacarelacs",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ReverseString(testCase.input, len(testCase.input)-1)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %s but got %s", (i + 1), testCase.expected, got)
			}
		})
	}
}
