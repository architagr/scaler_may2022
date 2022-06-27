package two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected []int
}

func TestTwoSum(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4, 5, 6},
		inputB:   6,
		expected: []int{2, 4},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{4, 7, -4, 2, 2, 2, 3, -5, -3, 9, -4, 9, -7, 7, -1, 9, 9, 4, 1, -4, -2, 3, -3, -5, 4, -7, 7, 9, -4, 4, -8},
		inputB:   -3,
		expected: []int{4, 8},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := TwoSum(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
