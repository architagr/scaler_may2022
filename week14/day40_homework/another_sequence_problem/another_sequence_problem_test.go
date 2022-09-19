package another_sequence_problem

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestGetValueAnotherSeq(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    3,
		expected: 7,
	})
	testCases = append(testCases, TestCase{
		input:    2,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := GetValueAnotherSeq(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
