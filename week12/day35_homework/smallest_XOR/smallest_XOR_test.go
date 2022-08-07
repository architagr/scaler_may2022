package smallest_XOR

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected int
}

func TestSmallestXor(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        3,
		b:        3,
		expected: 7,
	})

	testCases = append(testCases, TestCase{
		a:        15,
		b:        2,
		expected: 12,
	})
	testCases = append(testCases, TestCase{
		a:        15,
		b:        4,
		expected: 15,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := SmallestXor(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
