package combination_sum_ii

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	b        int
	expected [][]int
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{10, 1, 2, 7, 6, 1, 5},
		b:        8,
		expected: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 3, 1},
		b:        3,
		expected: [][]int{{1, 2}, {3}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := combinationSum(testCase.input, testCase.b)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
