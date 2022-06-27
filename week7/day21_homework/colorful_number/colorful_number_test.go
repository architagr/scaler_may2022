package colorful_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestCheckColorfulNumber(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    23,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    236,
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		input:    12345,
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CheckColorfulNumber(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
