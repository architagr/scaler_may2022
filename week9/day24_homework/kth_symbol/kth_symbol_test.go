package kth_symbol

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB int
	expected       int
}

func TestKthSymbol(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   2,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   1,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		inputA:   10,
		inputB:   3,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := KthSymbol(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
