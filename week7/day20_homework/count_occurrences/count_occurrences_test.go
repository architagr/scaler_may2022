package count_occurrences

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestCountOccurrence(t *testing.T) {

	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	input:    "abobc",
	// 	expected: 1,
	// })
	testCases = append(testCases, TestCase{
		input:    "bobob",
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CountOccurrence(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
