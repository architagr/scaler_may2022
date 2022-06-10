package reverse_bits

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    uint32
	expected uint32
}

func TestReverseBits(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    0,
		expected: 0,
	})

	testcases = append(testcases, TestCase{
		input:    3,
		expected: 3221225472,
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := reverseBits(testCase.input)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
