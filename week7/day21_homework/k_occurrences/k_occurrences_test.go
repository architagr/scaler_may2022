package k_occurrences

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB int
	inputC         []int
	expected       int
}

func TestGetSumKOccurrences(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   5,
		inputB:   2,
		inputC:   []int{1, 2, 2, 3, 3},
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		inputA:   5,
		inputB:   5,
		inputC:   []int{1, 2, 3, 4, 5},
		expected: -1,
	})
	testCases = append(testCases, TestCase{
		inputA:   6,
		inputB:   2,
		inputC:   []int{1000000000, 1000000000, 999999999, 999999999, 999999998, 1},
		expected: 999999992,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := GetSumKOccurrences(testCase.inputA, testCase.inputB, testCase.inputC)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
