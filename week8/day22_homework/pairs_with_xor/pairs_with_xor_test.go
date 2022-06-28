package pairs_with_xor

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected int
}

func TestPairsWithXor(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{5, 4, 10, 15, 7, 6},
		inputB:   5,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{3, 6, 8, 10, 15, 50},
		inputB:   5,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := PairsWithXor(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
