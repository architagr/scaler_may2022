package lcm

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected int
}

func TestFindLcm(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   3,
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		inputA:   9,
		inputB:   6,
		expected: 18,
	})
	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   6,
		expected: 6,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d and %d", testCase.inputA, testCase.inputB), func(tb *testing.T) {
			got := FindLcm(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d and %d expected %d but got %d", testCase.inputA, testCase.inputB, testCase.expected, got)
			}
		})
	}

}
