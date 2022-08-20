package pubg

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestGcd(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        []int{6, 4},
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Pubg(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
