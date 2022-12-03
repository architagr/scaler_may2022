package unbounded_knapsack

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB []int
	inputC         int
	expected       int
}

func TestUnboundedKnapSack(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{5},
		inputB:   []int{10},
		inputC:   10,
		expected: 5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := UnboundedKnapSack(testCase.inputC, testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
