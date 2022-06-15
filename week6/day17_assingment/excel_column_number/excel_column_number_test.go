package excel_column_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestGetCoumnNumber(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "AB",
		expected: 28,
	})
	testCases = append(testCases, TestCase{
		input:    "ABCD",
		expected: 19010,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %s", testCase.input), func(tb *testing.T) {
			got := GetCoumnNumber(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %s expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
