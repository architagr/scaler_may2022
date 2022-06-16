package leap_year

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestFindLeapYear(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    2020,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    2019,
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", testCase.input), func(tb *testing.T) {
			got := FindLeapYear(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}

}
