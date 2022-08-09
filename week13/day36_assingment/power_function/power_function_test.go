package power_function

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, c, expected int
}

func TestPow(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        2,
		b:        3,
		c:        3,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		a:        -2,
		b:        3,
		c:        3,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := Pow(testCase.a, testCase.b, testCase.c)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
