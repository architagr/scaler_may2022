package power_function

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB, inputC int
	expected               int
}

func TestPow(t *testing.T) {

	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	inputA:   2,
	// 	inputB:   3,
	// 	inputC:   3,
	// 	expected: 2,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   0,
	// 	inputB:   0,
	// 	inputC:   3,
	// 	expected: 0,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   -1,
	// 	inputB:   1,
	// 	inputC:   3,
	// 	expected: 2,
	// })

	testCases = append(testCases, TestCase{
		inputA:   71045970,
		inputB:   41535484,
		inputC:   64735492,
		expected: 20805472,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Pow(testCase.inputA, testCase.inputB, testCase.inputC)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
