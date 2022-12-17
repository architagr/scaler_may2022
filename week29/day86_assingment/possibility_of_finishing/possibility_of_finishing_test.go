package possibility_of_finishing

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA         int
	inputB, inputC []int
	expected       int
}

func TestPosibilityOfFinishing(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   3,
		inputB:   []int{1, 2},
		inputC:   []int{2, 3},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputA:   2,
		inputB:   []int{1, 2},
		inputC:   []int{2, 1},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := PosibilityOfFinishing(testCase.inputA, testCase.inputB, testCase.inputC)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
