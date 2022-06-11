package unset_bits

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int64
	inputB          int
}

func TestUnsetBits(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    93,
		inputB:   4,
		expected: 80,
	})

	testcases = append(testcases, TestCase{
		input:    25,
		inputB:   3,
		expected: 24,
	})

	testcases = append(testcases, TestCase{
		input:    37,
		inputB:   3,
		expected: 32,
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := UnsetBits(testCase.input, testCase.inputB)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
