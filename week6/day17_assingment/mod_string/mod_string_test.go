package mod_string

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	inputB   int
	expected int
}

func TestFindMod(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	input:    "143",
	// 	inputB:   2,
	// 	expected: 1,
	// })
	testCases = append(testCases, TestCase{
		input:    "43535321",
		inputB:   47,
		expected: 20,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %s", testCase.input), func(tb *testing.T) {
			got := findMod(testCase.input, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %s expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
