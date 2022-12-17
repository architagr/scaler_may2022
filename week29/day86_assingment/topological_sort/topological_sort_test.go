package topological_sort

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   int
	inputB   [][]int
	expected []int
}

func TestTopologicalSort(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   6,
		inputB:   [][]int{{6, 3}, {6, 1}, {5, 1}, {5, 2}, {3, 4}, {4, 2}},
		expected: []int{5, 6, 1, 3, 4, 2},
	})

	testCases = append(testCases, TestCase{
		inputA:   3,
		inputB:   [][]int{{1, 2}, {2, 3}, {3, 1}},
		expected: []int{},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := TopologicalSort(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
