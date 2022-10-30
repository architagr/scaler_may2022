package maximum_frequency_stack

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected []int
}

func TestMaxFrequencyStack(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 5},
			{1, 7},
			{1, 5},
			{1, 7},
			{1, 4},
			{1, 5},
			{2, 0},
			{2, 0},
			{2, 0},
			{2, 0},
		},
		expected: []int{-1, -1, -1, -1, -1, -1, 5, 7, 5, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MaxFrequencyStack(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
