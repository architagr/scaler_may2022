package distance_of_nearest_cell

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]int
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 1, 0}},
		expected: [][]int{{3, 2, 1, 0}, {2, 1, 0, 0}, {1, 0, 0, 1}},
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		expected: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
