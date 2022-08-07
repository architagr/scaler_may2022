package reverse_bits

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, expected uint32
}

func TestReverseBits(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        0,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        3,
		expected: 3221225472,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := ReverseBits(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
