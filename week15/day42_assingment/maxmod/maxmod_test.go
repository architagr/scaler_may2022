package maxmod

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestMaxMod(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{1, 2, 44, 3},
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		a:        []int{1, 2, 3, 3},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		a:        []int{5, 5, 5, 5},
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		a:        []int{927, 707, 374, 394, 279, 799, 878, 937, 431, 112},
		expected: 927,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MaxMod(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
