package subset_ii

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected [][]int
}

func TestSubsets2(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{9},
		expected: [][]int{{}, {9}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := subsets(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
