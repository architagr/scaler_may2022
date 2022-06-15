package divisibility_by_8

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestDivisibleBy8(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "16",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "123",
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %s", testCase.input), func(tb *testing.T) {
			got := DivisibleBy8(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %s expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
