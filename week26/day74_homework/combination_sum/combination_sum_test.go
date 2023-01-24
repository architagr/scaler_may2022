package combination_sum

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
		input:    []int{2, 3},
		b:        2,
		expected: [][]int{{2}},
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 3, 6, 7},
		b:        7,
		expected: [][]int{{2, 2, 3}, {7}},
	})
	testCases = append(testCases, TestCase{
		input: []int{8, 10, 6, 11, 1, 16, 8},
		b:     28,
		expected: [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 6},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 8},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 10},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 8},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 16},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 11},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6, 6, 6},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 10},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 11},
			{1, 1, 1, 1, 1, 1, 1, 1, 6, 6, 8},
			{1, 1, 1, 1, 1, 1, 1, 1, 10, 10},
			{1, 1, 1, 1, 1, 1, 1, 10, 11},
			{1, 1, 1, 1, 1, 1, 6, 6, 10},
			{1, 1, 1, 1, 1, 1, 6, 8, 8},
			{1, 1, 1, 1, 1, 1, 6, 16},
			{1, 1, 1, 1, 1, 1, 11, 11},
			{1, 1, 1, 1, 1, 6, 6, 11},
			{1, 1, 1, 1, 6, 6, 6, 6},
			{1, 1, 1, 1, 6, 8, 10},
			{1, 1, 1, 1, 8, 8, 8},
			{1, 1, 1, 1, 8, 16},
			{1, 1, 1, 6, 8, 11},
			{1, 1, 6, 6, 6, 8},
			{1, 1, 6, 10, 10},
			{1, 1, 8, 8, 10},
			{1, 1, 10, 16},
			{1, 6, 10, 11},
			{1, 8, 8, 11},
			{1, 11, 16},
			{6, 6, 6, 10},
			{6, 6, 8, 8},
			{6, 6, 16},
			{6, 11, 11},
			{8, 10, 10},
		},
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
