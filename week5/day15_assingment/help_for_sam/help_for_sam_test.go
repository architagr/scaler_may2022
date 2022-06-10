package help_for_sam

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestFindHelpTakenFromSame(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    5,
		expected: 2,
	})

	testcases = append(testcases, TestCase{
		input:    3,
		expected: 2,
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := FindHelpTakenFromSame(testCase.input)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
