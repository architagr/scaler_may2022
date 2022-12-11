package sum_of_min_and_max

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	inputB   int
	expected int
}

func TestMaxFrequencyStack(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{2, 5, -1, 7, -3, -1, -2},
		inputB:   4,
		expected: 18,
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, -1, 3},
		inputB:   2,
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := SumMaxMin(testCase.input, testCase.inputB)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
