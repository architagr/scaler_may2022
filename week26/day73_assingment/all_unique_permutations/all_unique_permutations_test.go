package all_unique_permutations

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected [][]int
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3},
		expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 2},
		expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := permute(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
