package chocolate_distribution

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a           []int
	b, expected int
}

func TestChocolateDistribution(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{3, 4, 1, 9, 56, 7, 9, 12},
		b:        5,
		expected: 6,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ChocolateDistribution(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
