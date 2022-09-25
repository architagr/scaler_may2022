package period_of_a_string

import (
	"fmt"
	"testing"
)

func TestPeriodOfString(t *testing.T) {
	type TestCase struct {
		input    string
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "abababab",
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    "aaaa",
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    "umeaylnlfdumeaylnlfdumeaylnlfd",
		expected: 10,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := PeriodOfString(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
