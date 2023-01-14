package uniqie_paths_3

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestUniquePaths3(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 1}, {2, 0}},
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := UniquePaths3(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
