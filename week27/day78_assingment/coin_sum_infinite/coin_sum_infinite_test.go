package coin_sum_infinite

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	b        int
	expected int
}

func TestCoinSumInfinite(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3},
		b:        4,
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		input:    []int{10},
		b:        10,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{18, 24, 23, 10, 16, 19, 2, 9, 5, 12, 17, 3, 28, 29, 4, 13, 15, 8},
		b:        458,
		expected: 867621,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CoinSumInfinite(testCase.input, testCase.b)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
