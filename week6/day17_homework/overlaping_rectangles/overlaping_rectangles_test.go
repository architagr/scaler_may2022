package overlaping_rectangles

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB, inputC, inputD, inputE, inputF, inputG, inputH, expected int
}

func TestFindOverlapingRectangles(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA: 0, inputB: 0,
		inputC: 4, inputD: 4,
		inputE: 2, inputF: 2,
		inputG: 6, inputH: 6,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA: 0, inputB: 0,
		inputC: 4, inputD: 4,
		inputE: 2, inputF: 2,
		inputG: 3, inputH: 3,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA: 12, inputB: 47,
		inputC: 39, inputD: 72,
		inputE: 8, inputF: 3,
		inputG: 47, inputH: 55,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA: 60, inputB: 65,
		inputC: 99, inputD: 84,
		inputE: 85, inputF: 24,
		inputG: 99, inputH: 84,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA: 0, inputB: 0,
		inputC: 1, inputD: 1,
		inputE: 1, inputF: 1,
		inputG: 6, inputH: 6,
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing # %d", (i+1)), func(tb *testing.T) {
			got := FindOverlapingRectangles(testCase.inputA, testCase.inputB, testCase.inputC, testCase.inputD, testCase.inputE, testCase.inputF, testCase.inputG, testCase.inputH)
			if got != testCase.expected {
				tb.Errorf("tested # %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
