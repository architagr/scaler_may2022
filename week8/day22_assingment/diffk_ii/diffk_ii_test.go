package diffk_ii

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected int
}

func TestDiffk2(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 5, 3},
		inputB:   2,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{11, 85, 100, 44, 3, 32, 96, 72, 93, 76, 67, 93, 63, 5, 10, 45, 99, 35, 13},
		inputB:   60,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Diffk2(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
