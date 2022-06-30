package sum_of_digits

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestSumOfDigits(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    123,
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		input:    4567,
		expected: 22,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := SumOfDigits(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
