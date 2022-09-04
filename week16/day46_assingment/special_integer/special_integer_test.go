package special_integer

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA           []int
	inputB, expected int
}

func TestFindSpecialInt(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4, 5},
		inputB:   10,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{5, 17, 100, 11},
		inputB:   130,
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{5, 10, 20, 100, 105},
		inputB:   130,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindSpecialInt(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
