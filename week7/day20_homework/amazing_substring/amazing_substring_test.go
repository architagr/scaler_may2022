package amazing_substring

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestFindAmazingSubstring(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "ABEC",
		expected: 6,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindAmazingSubstring(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
