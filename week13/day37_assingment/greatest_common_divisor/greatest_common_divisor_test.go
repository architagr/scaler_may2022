package greatest_common_divisor

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected int
}

func TestGcd(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        2,
		b:        3,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := Gcd(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
