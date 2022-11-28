package cuttingarod

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected int
}

func TestSolve(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{3, 4, 1, 6, 2},
		expected: 15,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 5, 2, 5, 6},
		expected: 11,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CuttingARod(testCase.inputA)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
