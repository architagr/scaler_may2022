package nth_magic_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestFindNthMagicNumber(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    3,
		expected: 30,
	})
	// testCases = append(testCases, TestCase{
	// 	input:    10,
	// 	expected: 650,
	// })

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", testCase.input), func(tb *testing.T) {
			got := FindNthMagicNumber(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}

}
