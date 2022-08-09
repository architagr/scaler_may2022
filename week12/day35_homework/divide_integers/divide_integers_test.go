package divide_integers

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected int
}

func TestDivide(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        10,
		b:        3,
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		a:        9,
		b:        3,
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		a:        8,
		b:        3,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		a:        -8,
		b:        3,
		expected: -2,
	})
	testCases = append(testCases, TestCase{
		a:        0,
		b:        3,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        2147483647,
		b:        1,
		expected: 2147483647,
	})
	testCases = append(testCases, TestCase{
		a:        -2147483648,
		b:        -1,
		expected: 2147483647,
	})
	testCases = append(testCases, TestCase{
		a:        -2147483648,
		b:        -10000000,
		expected: 214,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := Divide(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
