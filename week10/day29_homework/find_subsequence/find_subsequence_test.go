package find_subsequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB string
	expected       string
}

func TestSubSequence(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   "bit",
		inputB:   "dfbkjijgbbiihbmmt",
		expected: "YES",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := SubSequence(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
