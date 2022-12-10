package matrix_chain_multiplication

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected int
}

func TestMatrixChainMultiplication(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{40, 20, 30, 10, 30},
		expected: 26000,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{10, 20, 30},
		expected: 6000,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MatrixChainMultiplication(testCase.inputA)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
