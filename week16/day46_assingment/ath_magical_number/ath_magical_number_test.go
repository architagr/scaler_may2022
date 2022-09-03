package ath_magical_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB, inputC, expected int
}

func TestAthMagicalNumber(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	inputA:   1,
	// 	inputB:   2,
	// 	inputC:   3,
	// 	expected: 2,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   4,
	// 	inputB:   2,
	// 	inputC:   3,
	// 	expected: 6,
	// })
	// testCases = append(testCases, TestCase{
	// 	inputA:   6,
	// 	inputB:   2,
	// 	inputC:   3,
	// 	expected: 9,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   14,
	// 	inputB:   10,
	// 	inputC:   12,
	// 	expected: 84,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   807414236,
	// 	inputB:   3788,
	// 	inputC:   38141,
	// 	expected: 238134151,
	// })
	// testCases = append(testCases, TestCase{
	// 	inputA:   489039022,
	// 	inputB:   18923,
	// 	inputC:   32309,
	// 	expected: 119599631,
	// })
	testCases = append(testCases, TestCase{
		inputA:   827802315,
		inputB:   12423,
		inputC:   7916,
		expected: 678065970,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := AthMagicalNumber(testCase.inputA, testCase.inputB, testCase.inputC)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputB:   10,
		inputC:   12,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputB:   10,
		inputC:   25,
		expected: 5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := gcd(testCase.inputB, testCase.inputC)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
