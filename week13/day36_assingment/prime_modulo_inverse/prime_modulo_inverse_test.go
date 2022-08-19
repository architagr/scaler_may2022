package prime_modulo_inverse

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected int
}

func TestInverseModulo(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        3,
		b:        5,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		a:        6,
		b:        23,
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		a:        82993,
		b:        999996247,
		expected: 662764252,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := InverseModulo(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
