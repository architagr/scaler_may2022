package strange_equality

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, expected int
}

func TestStrangeEquality(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        5,
		expected: 10,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := StrangeEquality(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
